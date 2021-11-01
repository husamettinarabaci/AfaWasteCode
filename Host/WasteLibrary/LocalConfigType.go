package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//LocalConfigType
type LocalConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
	Key0       string
	Key1       string
	Key2       string
	Key3       string
	Key4       string
	Key5       string
	Key6       string
	Key7       string
	Key8       string
	Key9       string
	Key10      string
	Key11      string
	Key12      string
	Key13      string
	Key14      string
	Key15      string
	Key16      string
	Key17      string
	Key18      string
	Key19      string
	Key20      string
	Key21      string
	Key22      string
	Key23      string
	Key24      string
	Key25      string
	Key26      string
	Key27      string
	Key28      string
	Key29      string
	Key30      string
	Key31      string
	Key32      string
	Key33      string
	Key34      string
	Key35      string
	Key36      string
	Key37      string
	Key38      string
	Key39      string
	Key40      string
	Key41      string
	Key42      string
	Key43      string
	Key44      string
	Key45      string
	Key46      string
	Key47      string
	Key48      string
	Key49      string
	Key50      string
	Key51      string
	Key52      string
	Key53      string
	Key54      string
	Key55      string
	Key56      string
	Key57      string
	Key58      string
	Key59      string
	Key60      string
	Key61      string
	Key62      string
	Key63      string
	Key64      string
	Key65      string
	Key66      string
	Key67      string
	Key68      string
	Key69      string
	Key70      string
	Key71      string
	Key72      string
	Key73      string
	Key74      string
	Key75      string
	Key76      string
	Key77      string
	Key78      string
	Key79      string
	Key80      string
	Key81      string
	Key82      string
	Key83      string
	Key84      string
	Key85      string
	Key86      string
	Key87      string
	Key88      string
	Key89      string
	Key90      string
	Key91      string
	Key92      string
	Key93      string
	Key94      string
	Key95      string
	Key96      string
	Key97      string
	Key98      string
	Key99      string
	Key100     string
	Key101     string
	Key102     string
	Key103     string
	Key104     string
	Key105     string
	Key106     string
	Key107     string
	Key108     string
	Key109     string
	Key110     string
	Key111     string
	Key112     string
	Key113     string
	Key114     string
	Key115     string
	Key116     string
	Key117     string
	Key118     string
	Key119     string
	Key120     string
	Key121     string
	Key122     string
	Key123     string
	Key124     string
	Key125     string
	Key126     string
	Key127     string
	Key128     string
	Key129     string
	Key130     string
	Key131     string
	Key132     string
	Key133     string
	Key134     string
	Key135     string
	Key136     string
	Key137     string
	Key138     string
	Key139     string
	Key140     string
	Key141     string
	Key142     string
	Key143     string
	Key144     string
	Key145     string
	Key146     string
	Key147     string
	Key148     string
	Key149     string
	Key150     string
	Key151     string
	Key152     string
	Key153     string
	Key154     string
	Key155     string
	Key156     string
	Key157     string
	Key158     string
	Key159     string
	Key160     string
	Key161     string
	Key162     string
	Key163     string
	Key164     string
	Key165     string
	Key166     string
	Key167     string
	Key168     string
	Key169     string
	Key170     string
	Key171     string
	Key172     string
	Key173     string
	Key174     string
	Key175     string
	Key176     string
	Key177     string
	Key178     string
	Key179     string
	Key180     string
	Key181     string
	Key182     string
	Key183     string
	Key184     string
	Key185     string
	Key186     string
	Key187     string
	Key188     string
	Key189     string
	Key190     string
	Key191     string
	Key192     string
	Key193     string
	Key194     string
	Key195     string
	Key196     string
	Key197     string
	Key198     string
	Key199     string
	Key200     string
	Key201     string
	Key202     string
	Key203     string
	Key204     string
	Key205     string
	Key206     string
	Key207     string
	Key208     string
	Key209     string
	Key210     string
	Key211     string
	Key212     string
	Key213     string
	Key214     string
	Key215     string
	Key216     string
	Key217     string
	Key218     string
	Key219     string
	Key220     string
	Key221     string
	Key222     string
	Key223     string
	Key224     string
	Key225     string
	Key226     string
	Key227     string
	Key228     string
	Key229     string
	Key230     string
	Key231     string
	Key232     string
	Key233     string
	Key234     string
	Key235     string
	Key236     string
	Key237     string
	Key238     string
	Key239     string
	Key240     string
	Key241     string
	Key242     string
	Key243     string
	Key244     string
	Key245     string
	Key246     string
	Key247     string
	Key248     string
	Key249     string
	Key250     string
	Key251     string
	Key252     string
	Key253     string
	Key254     string
	Key255     string
	Key256     string
	Key257     string
	Key258     string
	Key259     string
	Key260     string
	Key261     string
	Key262     string
	Key263     string
	Key264     string
	Key265     string
	Key266     string
	Key267     string
	Key268     string
	Key269     string
	Key270     string
	Key271     string
	Key272     string
	Key273     string
	Key274     string
	Key275     string
	Key276     string
	Key277     string
	Key278     string
	Key279     string
	Key280     string
	Key281     string
	Key282     string
	Key283     string
	Key284     string
	Key285     string
	Key286     string
	Key287     string
	Key288     string
	Key289     string
	Key290     string
	Key291     string
	Key292     string
	Key293     string
	Key294     string
	Key295     string
	Key296     string
	Key297     string
	Key298     string
	Key299     string
}

//New
func (res *LocalConfigType) New() {
	res.CustomerId = 0
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()

	res.Key0 = ""
	res.Key1 = ""
	res.Key2 = ""
	res.Key3 = ""
	res.Key4 = ""
	res.Key5 = ""
	res.Key6 = ""
	res.Key7 = ""
	res.Key8 = ""
	res.Key9 = ""
	res.Key10 = ""
	res.Key11 = ""
	res.Key12 = ""
	res.Key13 = ""
	res.Key14 = ""
	res.Key15 = ""
	res.Key16 = ""
	res.Key17 = ""
	res.Key18 = ""
	res.Key19 = ""
	res.Key20 = ""
	res.Key21 = ""
	res.Key22 = ""
	res.Key23 = ""
	res.Key24 = ""
	res.Key25 = ""
	res.Key26 = ""
	res.Key27 = ""
	res.Key28 = ""
	res.Key29 = ""
	res.Key30 = ""
	res.Key31 = ""
	res.Key32 = ""
	res.Key33 = ""
	res.Key34 = ""
	res.Key35 = ""
	res.Key36 = ""
	res.Key37 = ""
	res.Key38 = ""
	res.Key39 = ""
	res.Key40 = ""
	res.Key41 = ""
	res.Key42 = ""
	res.Key43 = ""
	res.Key44 = ""
	res.Key45 = ""
	res.Key46 = ""
	res.Key47 = ""
	res.Key48 = ""
	res.Key49 = ""
	res.Key50 = ""
	res.Key51 = ""
	res.Key52 = ""
	res.Key53 = ""
	res.Key54 = ""
	res.Key55 = ""
	res.Key56 = ""
	res.Key57 = ""
	res.Key58 = ""
	res.Key59 = ""
	res.Key60 = ""
	res.Key61 = ""
	res.Key62 = ""
	res.Key63 = ""
	res.Key64 = ""
	res.Key65 = ""
	res.Key66 = ""
	res.Key67 = ""
	res.Key68 = ""
	res.Key69 = ""
	res.Key70 = ""
	res.Key71 = ""
	res.Key72 = ""
	res.Key73 = ""
	res.Key74 = ""
	res.Key75 = ""
	res.Key76 = ""
	res.Key77 = ""
	res.Key78 = ""
	res.Key79 = ""
	res.Key80 = ""
	res.Key81 = ""
	res.Key82 = ""
	res.Key83 = ""
	res.Key84 = ""
	res.Key85 = ""
	res.Key86 = ""
	res.Key87 = ""
	res.Key88 = ""
	res.Key89 = ""
	res.Key90 = ""
	res.Key91 = ""
	res.Key92 = ""
	res.Key93 = ""
	res.Key94 = ""
	res.Key95 = ""
	res.Key96 = ""
	res.Key97 = ""
	res.Key98 = ""
	res.Key99 = ""
	res.Key100 = ""
	res.Key101 = ""
	res.Key102 = ""
	res.Key103 = ""
	res.Key104 = ""
	res.Key105 = ""
	res.Key106 = ""
	res.Key107 = ""
	res.Key108 = ""
	res.Key109 = ""
	res.Key110 = ""
	res.Key111 = ""
	res.Key112 = ""
	res.Key113 = ""
	res.Key114 = ""
	res.Key115 = ""
	res.Key116 = ""
	res.Key117 = ""
	res.Key118 = ""
	res.Key119 = ""
	res.Key120 = ""
	res.Key121 = ""
	res.Key122 = ""
	res.Key123 = ""
	res.Key124 = ""
	res.Key125 = ""
	res.Key126 = ""
	res.Key127 = ""
	res.Key128 = ""
	res.Key129 = ""
	res.Key130 = ""
	res.Key131 = ""
	res.Key132 = ""
	res.Key133 = ""
	res.Key134 = ""
	res.Key135 = ""
	res.Key136 = ""
	res.Key137 = ""
	res.Key138 = ""
	res.Key139 = ""
	res.Key140 = ""
	res.Key141 = ""
	res.Key142 = ""
	res.Key143 = ""
	res.Key144 = ""
	res.Key145 = ""
	res.Key146 = ""
	res.Key147 = ""
	res.Key148 = ""
	res.Key149 = ""
	res.Key150 = ""
	res.Key151 = ""
	res.Key152 = ""
	res.Key153 = ""
	res.Key154 = ""
	res.Key155 = ""
	res.Key156 = ""
	res.Key157 = ""
	res.Key158 = ""
	res.Key159 = ""
	res.Key160 = ""
	res.Key161 = ""
	res.Key162 = ""
	res.Key163 = ""
	res.Key164 = ""
	res.Key165 = ""
	res.Key166 = ""
	res.Key167 = ""
	res.Key168 = ""
	res.Key169 = ""
	res.Key170 = ""
	res.Key171 = ""
	res.Key172 = ""
	res.Key173 = ""
	res.Key174 = ""
	res.Key175 = ""
	res.Key176 = ""
	res.Key177 = ""
	res.Key178 = ""
	res.Key179 = ""
	res.Key180 = ""
	res.Key181 = ""
	res.Key182 = ""
	res.Key183 = ""
	res.Key184 = ""
	res.Key185 = ""
	res.Key186 = ""
	res.Key187 = ""
	res.Key188 = ""
	res.Key189 = ""
	res.Key190 = ""
	res.Key191 = ""
	res.Key192 = ""
	res.Key193 = ""
	res.Key194 = ""
	res.Key195 = ""
	res.Key196 = ""
	res.Key197 = ""
	res.Key198 = ""
	res.Key199 = ""
	res.Key200 = ""
	res.Key201 = ""
	res.Key202 = ""
	res.Key203 = ""
	res.Key204 = ""
	res.Key205 = ""
	res.Key206 = ""
	res.Key207 = ""
	res.Key208 = ""
	res.Key209 = ""
	res.Key210 = ""
	res.Key211 = ""
	res.Key212 = ""
	res.Key213 = ""
	res.Key214 = ""
	res.Key215 = ""
	res.Key216 = ""
	res.Key217 = ""
	res.Key218 = ""
	res.Key219 = ""
	res.Key220 = ""
	res.Key221 = ""
	res.Key222 = ""
	res.Key223 = ""
	res.Key224 = ""
	res.Key225 = ""
	res.Key226 = ""
	res.Key227 = ""
	res.Key228 = ""
	res.Key229 = ""
	res.Key230 = ""
	res.Key231 = ""
	res.Key232 = ""
	res.Key233 = ""
	res.Key234 = ""
	res.Key235 = ""
	res.Key236 = ""
	res.Key237 = ""
	res.Key238 = ""
	res.Key239 = ""
	res.Key240 = ""
	res.Key241 = ""
	res.Key242 = ""
	res.Key243 = ""
	res.Key244 = ""
	res.Key245 = ""
	res.Key246 = ""
	res.Key247 = ""
	res.Key248 = ""
	res.Key249 = ""
	res.Key250 = ""
	res.Key251 = ""
	res.Key252 = ""
	res.Key253 = ""
	res.Key254 = ""
	res.Key255 = ""
	res.Key256 = ""
	res.Key257 = ""
	res.Key258 = ""
	res.Key259 = ""
	res.Key260 = ""
	res.Key261 = ""
	res.Key262 = ""
	res.Key263 = ""
	res.Key264 = ""
	res.Key265 = ""
	res.Key266 = ""
	res.Key267 = ""
	res.Key268 = ""
	res.Key269 = ""
	res.Key270 = ""
	res.Key271 = ""
	res.Key272 = ""
	res.Key273 = ""
	res.Key274 = ""
	res.Key275 = ""
	res.Key276 = ""
	res.Key277 = ""
	res.Key278 = ""
	res.Key279 = ""
	res.Key280 = ""
	res.Key281 = ""
	res.Key282 = ""
	res.Key283 = ""
	res.Key284 = ""
	res.Key285 = ""
	res.Key286 = ""
	res.Key287 = ""
	res.Key288 = ""
	res.Key289 = ""
	res.Key290 = ""
	res.Key291 = ""
	res.Key292 = ""
	res.Key293 = ""
	res.Key294 = ""
	res.Key295 = ""
	res.Key296 = ""
	res.Key297 = ""
	res.Key298 = ""
	res.Key299 = ""
}

//ToId String
func (res LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res LocalConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res LocalConfigType) ToString() string {
	return string(res.ToByte())

}

//Byte To LocalConfigType
func ByteToLocalConfigType(retByte []byte) LocalConfigType {
	var retVal LocalConfigType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To LocalConfigType
func StringToLocalConfigType(retStr string) LocalConfigType {
	return ByteToLocalConfigType([]byte(retStr))
}
