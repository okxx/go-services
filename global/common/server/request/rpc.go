package request

//rpc
type Rpc struct {
	Node 		map[string]string //节点

 	//服务集合，支持多服务调用  k=> services uri  v => service params, type map
 	//{"/user/common/getInfo" : map[string]interface{}{"a":1}}
	Services 	interface{}
}

//rpc sync call services
func (rpc *Rpc) Sync() {

}