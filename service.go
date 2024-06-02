package entpb

// type ServiceOption func(msg *message)

// func WithService() MessageOption {
// 	return func(msg *message) {
// 		msg.Service = &ServiceDescriptor{}
// 	}
// }

// func Rpc(name string, req PbType, res PbType) MessageOption {
// 	return func(msg *messageAnnotation) {
// 		if msg.service == nil {
// 			msg.service = &ServiceDescriptor{}
// 		}
// 		msg.service.Rpcs = append(msg.service.Rpcs, &RpcDescriptor{
// 			Name:    name,
// 			ReqType: req,
// 			ResType: res,
// 		})
// 	}
// }

// type ServiceDescriptor struct {
// 	Base *MessageDescriptor

// 	Name string
// 	Rpcs []*RpcDescriptor

// 	Comment string
// }

// type RpcDescriptor struct {
// 	Name    string
// 	ReqType PbType
// 	ResType PbType
// }
