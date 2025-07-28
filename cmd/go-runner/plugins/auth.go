package plugins

import (
	"encoding/json"
	"net/http"

	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
)

func init() {
	err := plugin.RegisterPlugin(&Auth{})
	if err != nil {
		log.Fatalf("failed to register plugin say: %s", err)
	}
}

// Auth is a demo to show how to return data directly instead of proxying
// it to the upstream.
type Auth struct {
	plugin.DefaultPlugin
}

type AuthConf struct {
	Body string `json:"body"`
}

// 此处必须实现三个方法Name,ParseConf,Filter
// 主要逻辑在Filter中

func (p *Auth) Name() string {
	return "auth"
}

func (p *Auth) ParseConf(in []byte) (interface{}, error) {
	conf := AuthConf{}
	err := json.Unmarshal(in, &conf)
	return conf, err
}

// 会在每个配置了 auth 插件的请求中执行
func (p *Auth) Filter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	// 如果认证通过(不对请求做拦截,直接转发服务),直接return即可,不要操作response对象
	// 如果未通过,则可以在response增加http状态码,返回body等拦截返回

	// 根据访问path做限制
	// 1. 如果path等于/login表示原封不动转发到其他服务
	if string(r.Path()) == "/login" {
		// 直接return即可转发到路由配置的upstream中
		return

		// 如果path等于/auth则表示当前接口需要认证
	} else if string(r.Path()) == "/auth" {
		// 通过http调用内部认证,成功则正常返回,认证失败则返回异常信息,此处示例为认证失败
		httpResponseCode := 409
		if httpResponseCode == 200 {
			// 如果认证接口返回200则认证通过,直接return转发接口
			return

			// 如果调用认证接口失败,则返回认证失败
		} else {
			// 不通过则构建响应对象,返回给前端
			body := `"code":"409","desc":"auth err"`
			// 自定义Header
			w.Header().Add("X-Resp-A6-Runner", "Go")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(409)
			_, err := w.Write([]byte(body))
			if err != nil {
				log.Errorf("failed to write: %s", err)
			}
			return
		}
	}
	// 其他接口则将path拦截,原封不动返回
	w.Header().Add("X-Resp-A6-Runner", "Go")
	_, err := w.Write(r.Path())
	if err != nil {
		log.Errorf("failed to write: %s", err)
	}
	return
}
