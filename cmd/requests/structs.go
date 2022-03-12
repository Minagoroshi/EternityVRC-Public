package requests

type EternityPost struct {
	WorldID    string `json:"worldID"`
	ActorID    string `json:"actorID"`
	AvatarName string `json:"avatarName"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	Method     string `json:"method"`
	Time       int    `json:"time"`
	Mode       string `json:"mode"`
	Username   string `json:"username"`
	Sha3       string `json:"Sh3-H"`
}
