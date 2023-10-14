package zachestnyibiznes

type Сontacts struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Body    struct {
		Total int        `json:"total,omitempty"`
		Docs  []DocsItem `json:"docs,omitempty"`
	} `json:"body,omitempty"`
}

type DocsItem struct {
	NAMING_FAILED   string `json:"РосстатОбщ,omitempty"`
	NAMING_FAILED0  string `json:"ЗакупкиОбщ,omitempty"`
	NAMING_FAILED1  string `json:"РосТрудОбщ,omitempty"`
	NAMING_FAILED2  string `json:"РКНСайтОбщ,omitempty"`
	NAMING_FAILED3  string `json:"РеестрАккредОбщ,omitempty"`
	NAMING_FAILED4  string `json:"РКНСмиОбщ,omitempty"`
	Num1            string `json:"Источник1Общ,omitempty"`
	Num2            string `json:"Источник2Общ,omitempty"`
	NAMING_FAILED5  string `json:"ПочтаРосРаб,omitempty"`
	NAMING_FAILED6  string `json:"ПочтРОССТАТ,omitempty"`
	NAMING_FAILED7  string `json:"ПочтПерсРКН,omitempty"`
	NAMING_FAILED8  string `json:"ПочтЗАКУПКИ,omitempty"`
	NAMING_FAILED9  string `json:"ТелРосРаб,omitempty"`
	NAMING_FAILED10 string `json:"ТелРОССТАТ,omitempty"`
	NAMING_FAILED11 string `json:"ТелПерсРКН,omitempty"`
	NAMING_FAILED12 string `json:"ТелЗАКУПКИ,omitempty"`
	NAMING_FAILED13 string `json:"Тел,omitempty"`
	NAMING_FAILED14 string `json:"НомТелФНС,omitempty"`
	NAMING_FAILED15 string `json:"ТелРеестрАккред,omitempty"`
	NAMING_FAILED16 string `json:"ПочтРеестрАккред,omitempty"`
	NAMING_FAILED17 string `json:"СайтРеестрАккред,omitempty"`
	NAMING_FAILED18 string `json:"ТелРКНСми,omitempty"`
	NAMING_FAILED19 string `json:"ПочтРКНСми,omitempty"`
	NAMING_FAILED20 string `json:"СайтРКНСми,omitempty"`
	Num10           string `json:"ТелИсточник1,omitempty"`
	Num11           string `json:"ПочтИсточник1,omitempty"`
	Num12           string `json:"СайтИсточник1,omitempty"`
	Num20           string `json:"ТелИсточник2,omitempty"`
	Num21           string `json:"ПочтИсточник2,omitempty"`
	Num22           string `json:"СайтИсточник2,omitempty"`
	Num3            string `json:"ТелИсточник3,omitempty"`
	Num30           string `json:"ПочтИсточник3,omitempty"`
	Num31           string `json:"СайтИсточник3,omitempty"`
	Num32           string `json:"АдресИсточник3,omitempty"`
	NAMING_FAILED21 string `json:"Адрес,omitempty"`
	NAMING_FAILED22 string `json:"ПризнОтсутАдресЮЛ,omitempty"`
	NAMING_FAILED23 string `json:"ТипУлица,omitempty"`
	NAMING_FAILED24 string `json:"НаимУлица,omitempty"`
	NAMING_FAILED25 string `json:"ТипРайон,omitempty"`
	NAMING_FAILED26 string `json:"НаимРайон,omitempty"`
	NAMING_FAILED27 string `json:"ТипНаселПункт,omitempty"`
	NAMING_FAILED28 string `json:"НаимНаселПункт,omitempty"`
	NAMING_FAILED29 string `json:"ТипГород,omitempty"`
	NAMING_FAILED30 string `json:"НаимГород,omitempty"`
	NAMING_FAILED31 string `json:"ТипРегион,omitempty"`
	NAMING_FAILED32 string `json:"НаимРегион,omitempty"`
	NAMING_FAILED33 string `json:"Дом,omitempty"`
	NAMING_FAILED34 string `json:"Корпус,omitempty"`
	NAMING_FAILED35 string `json:"Кварт,omitempty"`
	NAMING_FAILED36 string `json:"КодРегион,omitempty"`
	NAMING_FAILED37 string `json:"КодАдрКладр,omitempty"`
	NAMING_FAILED38 string `json:"Индекс,omitempty"`
	NAMING_FAILED39 bool   `json:"СвНедАдресЮЛ,omitempty"`
}
