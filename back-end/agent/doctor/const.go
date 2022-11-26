package doctor

type major_doctor int8

const (
	Obstetrics_And_Gynecology major_doctor = 0; // 妇产科
	Pediatrics major_doctor = 1; // 儿科
	Neurology major_doctor = 2; // 神经科
	Gastroenterology major_doctor = 3; // 消化科
	orthopedics major_doctor = 4; // 骨科
	Respiratory major_doctor = 5; // 呼吸科
	ENT major_doctor = 6; // 耳鼻喉科
)