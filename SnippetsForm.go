package main

//GetSnipXML снипеты для генерации форм
func GetSnipXML() map[string]string {
	m := make(map[string]string)
	m["stylelabel"] ="" 
	m["ref"] = "<part class='label'>"+
	"<style>"+
		"<property name='Geometry'>{12,20,59,24}</property>"+
		"<property name='Caption'>[Код]</property>"+
		"<property name='TextColor'>2236962</property>"+
		"<property name='BackColor'>12632256</property>"+
		"<property name='Transparent'>.T.</property>"+
	"</style>"+
"</part>"+
"<part class='editbox'>"+
	"<style>"+
		"<property name='Geometry'>{228,20,294,32}</property>"+
		"<property name='Anchor'>130</property>"+
		"<property name='BackColor'>15395562</property>"+
		"<property name='nMaxLength'>48</property>"+
		"<property name='varName'>[m_title]</property>"+
	"</style>"+
"</part>";

	m["styleform"] = `<style>
	<property name='Geometry'>{479,185,529,159}</property>
	<property name="Name">[RefElem]</property>
	<property name="Caption">[Редактор элемента справочника]</property>
	<property name="Font">
		<font name="Georgia" width="0" height="-20M6.22" weight="400" charset="204"/>
	</property>
	<property name="FormType">[dlgModal]</property>
	<property name="lClipper">.F.</property>
	<property name="lExitOnEnter">.F.</property>
	<property name="lDebug">.F.</property>
	<property name="Variables">{nBox}</property>
	<property name="FromStyle">[Popup]</property>
	<property name="Icon">[False]</property>
	<property name="BackColor">12632256</property>
	<property name="NoModal">.F.</property>
	<property name="SystemMenu">.F.</property>
	<property name="Minimizebox">.F.</property>
	<property name="Maximizebox">.F.</property>
	<property name="AbsAlignent">.T.</property>
	<property name="SizeBox">.T.</property>
	<property name="Visible">.T.</property>
	<property name="3DLook">.F.</property>
	<property name="Clipsiblings">.F.</property>
	<property name="Clipchildren">.F.</property>
</style>
`
	return m
}
