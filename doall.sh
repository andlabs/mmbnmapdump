# 4 september 2013
awk '
BEGIN {
	print "mkdir map"
	command = "./mmbnmapdump ~/ida/MMBN\\ USA.gba"
}
!/LevelBackground </ { next }
$2 ~ /LB/ {
	gsub(/LB_/, "")
	area = $2
	subarea = 0
	# fall through
}
/<0>/ {		# invalid
	printf "true && # %s:%X invalid\n", area, subarea
	subarea++
	next
}
{			# valid
	gsub(/ROM:08/, "")
	printf "%s %s > map/%s_%X.png &&\n", command, $1, area, subarea
	subarea++
}
END { print "true" }
' <<\END
ROM:08012C54 LB_0            LevelBackground <unk_83B0298, unk_83B3230, unk_83B33D4>
ROM:08012C54                                         ; DATA XREF: ROM:off_8012BF4o
ROM:08012C54                                         ; 0
ROM:08012C60                 LevelBackground <unk_83B0298, unk_83B3230, unk_83B4054> ; 1
ROM:08012C6C                 LevelBackground <unk_83B4E98, unk_83B4CF4, unk_83B7980> ; 2
ROM:08012C78                 LevelBackground <unk_83B9148, unk_83BBEC0, unk_83BC064> ; 3
ROM:08012C84                 LevelBackground <0>     ; $
ROM:08012C90                 LevelBackground <unk_83B0298, unk_83BF0B4, unk_83BF258> ; 5
ROM:08012C9C                 LevelBackground <unk_83B0298, unk_83BFEF0, unk_83C0094> ; 6
ROM:08012CA8                 LevelBackground <unk_83B0298, unk_83C0D04, unk_83C0EA8> ; 7
ROM:08012CB4                 LevelBackground <unk_83C1CA8, unk_83C434C, unk_83C44F0> ; 8
ROM:08012CC0                 LevelBackground <unk_83B9148, unk_83BBEC0, unk_83C51D4> ; 9
ROM:08012CCC                 LevelBackground <0>     ; $A
ROM:08012CD8                 LevelBackground <unk_83B9148, unk_83C8280, unk_83C8424> ; $B
ROM:08012CE4                 LevelBackground <unk_83C8F94, unk_83CBF18, unk_83CC0BC> ; $C
ROM:08012CF0                 LevelBackground <unk_83CD220, unk_83CFBD0, unk_83CFD74> ; $D
ROM:08012CFC                 LevelBackground <unk_83B9148, unk_83D10CC, unk_83D1270> ; $E
ROM:08012D08 LB_1            LevelBackground <unk_83D22F0, unk_83D8104, unk_83D82A8>
ROM:08012D08                                         ; DATA XREF: ROM:08012BF8o
ROM:08012D08                                         ; 0
ROM:08012D14                 LevelBackground <unk_83DD8FC, unk_83E1FC0, unk_83E2164> ; 1
ROM:08012D20                 LevelBackground <unk_83E3CC4, unk_83E6678, unk_83E681C> ; 2
ROM:08012D2C                 LevelBackground <unk_83E7718, unk_83E9500, unk_83E96A4> ; 3 - Lan's room
ROM:08012D38                 LevelBackground <0>     ; 4
ROM:08012D44                 LevelBackground <unk_83EA180, unk_83EC5F0, unk_83EC794> ; 5
ROM:08012D50                 LevelBackground <unk_83ED76C, unk_83EF004, unk_83EF1A8> ; 6
ROM:08012D5C                 LevelBackground <unk_83EFBD4, unk_83F2310, unk_83F24B4> ; 7
ROM:08012D68                 LevelBackground <0>     ; 8
ROM:08012D74                 LevelBackground <unk_83F308C, unk_83F600C, unk_83F61B0> ; 9
ROM:08012D80                 LevelBackground <0>     ; $A
ROM:08012D8C                 LevelBackground <unk_83F715C, unk_83F9D04, unk_83F9EA8> ; $B
ROM:08012D98                 LevelBackground <unk_83FACE4, unk_83FEE84, unk_83FF028> ; $C
ROM:08012DA4                 LevelBackground <unk_83FACE4, unk_84005B0, unk_8400754> ; $D
ROM:08012DB0 LB_2            LevelBackground <unk_8401ABC, unk_8405374, unk_8405518>
ROM:08012DB0                                         ; DATA XREF: ROM:08012BFCo
ROM:08012DB0                                         ; 0
ROM:08012DBC                 LevelBackground <unk_83FACE4, unk_840A4A0, unk_840A644> ; 1
ROM:08012DC8                 LevelBackground <unk_840BDD0, unk_840E520, unk_840E6C4> ; 2
ROM:08012DD4                 LevelBackground <unk_8410FA4, unk_841338C, unk_8413530> ; 3
ROM:08012DE0                 LevelBackground <unk_8415D74, unk_841640C, unk_84165B0> ; 4
ROM:08012DEC                 LevelBackground <unk_84176B8, unk_8419D04, unk_8419EA8> ; 5
ROM:08012DF8                 LevelBackground <unk_841C27C, unk_841D8F0, unk_841DA94> ; 6
ROM:08012E04                 LevelBackground <unk_841F390, unk_8420B50, unk_8420CF4> ; 7
ROM:08012E10                 LevelBackground <0>     ; 8
ROM:08012E1C                 LevelBackground <unk_842344C, unk_8426324, unk_84264C8> ; 9
ROM:08012E28                 LevelBackground <0>     ; $A
ROM:08012E34                 LevelBackground <unk_842344C, unk_8428498, unk_842863C> ; $B
ROM:08012E40 LB_3            LevelBackground <unk_842BDB8, unk_842ED5C, unk_842EF00>
ROM:08012E40                                         ; DATA XREF: ROM:08012C00o
ROM:08012E40                                         ; 0
ROM:08012E4C                 LevelBackground <unk_83FACE4, unk_8432970, unk_8432B14> ; 1
ROM:08012E58                 LevelBackground <unk_843423C, unk_8438430, unk_84385D4> ; 2
ROM:08012E64                 LevelBackground <unk_843423C, unk_843A984, unk_843AB28> ; 3
ROM:08012E70                 LevelBackground <unk_843423C, unk_843CE70, unk_843D014> ; 4
ROM:08012E7C                 LevelBackground <unk_843423C, unk_843F534, unk_843F6D8> ; 5
ROM:08012E88                 LevelBackground <unk_8441810, unk_8443618, unk_84437BC> ; 6
ROM:08012E94                 LevelBackground <unk_84442B0, unk_844531C, unk_84454C0> ; 7
ROM:08012EA0 LB_4            LevelBackground <unk_8445F98, unk_8447BB8, unk_8447D5C>
ROM:08012EA0                                         ; DATA XREF: ROM:08012C04o
ROM:08012EA0                                         ; 0
ROM:08012EAC                 LevelBackground <unk_844A1C4, unk_844CB4C, unk_844CCF0> ; 1
ROM:08012EB8                 LevelBackground <unk_8445F98, unk_844F584, unk_844F728> ; 2
ROM:08012EC4                 LevelBackground <unk_8445F98, unk_8451AFC, unk_8451CA0> ; 3
ROM:08012ED0                 LevelBackground <unk_841F390, unk_84537D8, unk_845397C> ; 4
ROM:08012EDC                 LevelBackground <unk_8454EEC, unk_845673C, unk_84568E0> ; 5
ROM:08012EE8 LB_5            LevelBackground <unk_8458D1C, unk_845C3EC, unk_845C590>
ROM:08012EE8                                         ; DATA XREF: ROM:08012C08o
ROM:08012EE8                                         ; 0
ROM:08012EF4                 LevelBackground <unk_845E93C, unk_8462CC4, unk_8462E68> ; 1
ROM:08012F00                 LevelBackground <unk_8465260, unk_8468E34, unk_8468FD8> ; 2
ROM:08012F0C                 LevelBackground <unk_846B200, unk_846C410, unk_846C5B4> ; 3
ROM:08012F18                 LevelBackground <unk_846B200, unk_846D578, unk_846D71C> ; 4
ROM:08012F24                 LevelBackground <unk_846B200, unk_846E6B0, unk_846E854> ; 5
ROM:08012F30 LB_80           LevelBackground <unk_846F8A8, unk_8472B84, unk_8472D28>
ROM:08012F30                                         ; DATA XREF: ROM:off_8012C0Co
ROM:08012F30                                         ; 0
ROM:08012F3C                 LevelBackground <unk_8475A74, unk_8478820, unk_84789C4> ; 1
ROM:08012F48                 LevelBackground <unk_846F8A8, unk_847B634, unk_847B7D8> ; 2
ROM:08012F54                 LevelBackground <unk_846F8A8, unk_847E290, unk_847E434> ; 3
ROM:08012F60                 LevelBackground <unk_846F8A8, unk_8481450, unk_84815F4> ; 4
ROM:08012F6C LB_81           LevelBackground <unk_8483F78, unk_8487718, unk_84878BC>
ROM:08012F6C                                         ; DATA XREF: ROM:08012C10o
ROM:08012F6C                                         ; 0
ROM:08012F78                 LevelBackground <unk_8483F78, unk_8487718, unk_848BB6C> ; 1
ROM:08012F84 LB_82           LevelBackground <unk_848EE68, unk_84955C4, unk_8495768>
ROM:08012F84                                         ; DATA XREF: ROM:08012C14o
ROM:08012F84                                         ; 0
ROM:08012F90                 LevelBackground <unk_848EE68, unk_84955C4, unk_849A91C> ; 1
ROM:08012F9C                 LevelBackground <unk_848EE68, unk_84955C4, unk_849EC7C> ; 2
ROM:08012FA8                 LevelBackground <unk_84A4264, unk_84A99D0, unk_84A9B74> ; 3
ROM:08012FB4                 LevelBackground <unk_84AF050, unk_84B26C4, unk_84B2868> ; 4
ROM:08012FC0                 LevelBackground <unk_84A4264, unk_84B6EE0, unk_84B7084> ; 5
ROM:08012FCC LB_83           LevelBackground <unk_84BBFB0, unk_84BDB30, unk_84BDCD4>
ROM:08012FCC                                         ; DATA XREF: ROM:08012C18o
ROM:08012FCC                                         ; 0
ROM:08012FD8                 LevelBackground <unk_84BBFB0, unk_84BDB30, unk_84C16F0> ; 1
ROM:08012FE4                 LevelBackground <unk_84BBFB0, unk_84BDB30, unk_84C5284> ; 2
ROM:08012FF0                 LevelBackground <unk_84BBFB0, unk_84BDB30, unk_84C9438> ; 3
ROM:08012FFC                 LevelBackground <unk_84BBFB0, unk_84BDB30, unk_84CCC70> ; 4
ROM:08013008 LB_84           LevelBackground <unk_84D16FC, unk_84D3868, unk_84D3A0C>
ROM:08013008                                         ; DATA XREF: ROM:08012C1Co
ROM:08013008                                         ; 0
ROM:08013014                 LevelBackground <unk_84D16FC, unk_84D3868, unk_84D8530> ; 1
ROM:08013020                 LevelBackground <unk_84D16FC, unk_84D3868, unk_84DCF48> ; 2
ROM:0801302C                 LevelBackground <unk_84D16FC, unk_84D3868, unk_84E1AE8> ; 3
ROM:08013038 LB_85           LevelBackground <unk_8483F78, unk_8487718, unk_84E63AC>
ROM:08013038                                         ; DATA XREF: ROM:08012C20o
ROM:08013038                                         ; 0
ROM:08013044                 LevelBackground <unk_8475A74, unk_84EAFBC, unk_84EB160> ; 1
ROM:08013050                 LevelBackground <unk_848EE68, unk_84955C4, unk_84EEFC0> ; 2
ROM:0801305C                 LevelBackground <unk_84BBFB0, unk_84BDB30, unk_84F3DE0> ; 3
ROM:08013068                 LevelBackground <unk_84D16FC, unk_84D3868, unk_84F8438> ; 4
ROM:08013074                 LevelBackground <unk_84FCBA4, unk_84FDCFC, unk_84FDEA0> ; 5
ROM:08013080 LB_88           LevelBackground <unk_8500FD4, unk_8502C60, unk_8502E04>
ROM:08013080                                         ; DATA XREF: ROM:08012C2Co
ROM:08013080                                         ; 0
ROM:0801308C                 LevelBackground <unk_8500FD4, unk_8506384, unk_8503934> ; 1
ROM:08013098                 LevelBackground <unk_8500FD4, unk_85087D4, unk_8506528> ; 2
ROM:080130A4                 LevelBackground <unk_8500FD4, unk_850C900, unk_8508978> ; 3
ROM:080130B0 LB_89           LevelBackground <unk_850CAA4, unk_850DC6C, unk_850DE10>
ROM:080130B0                                         ; DATA XREF: ROM:08012C30o
ROM:080130B0                                         ; 0
ROM:080130BC                 LevelBackground <unk_850CAA4, unk_8511580, unk_850F498> ; 1
ROM:080130C8 LB_8A           LevelBackground <unk_850CAA4, unk_851398C, unk_8511724>
ROM:080130C8                                         ; DATA XREF: ROM:08012C34o
ROM:080130C8                                         ; 0
ROM:080130D4 LB_8B           LevelBackground <unk_850CAA4, unk_8517C78, unk_8513B30>
ROM:080130D4                                         ; DATA XREF: ROM:08012C38o
ROM:080130D4                                         ; 0
ROM:080130E0 LB_8C           LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C>
ROM:080130E0                                         ; DATA XREF: ROM:08012C3Co
ROM:080130E0                                         ; 0
ROM:080130EC                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 1
ROM:080130F8                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 2
ROM:08013104                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 3
ROM:08013110                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 4
ROM:0801311C                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 5
ROM:08013128                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 6
ROM:08013134                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 7
ROM:08013140                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 8
ROM:0801314C                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; 9
ROM:08013158                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; $A
ROM:08013164                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; $B
ROM:08013170                 LevelBackground <unk_850CAA4, unk_851A1C4, unk_8517E1C> ; $C
ROM:0801317C LB_90           LevelBackground <unk_851A368, unk_851E23C, unk_851E3E0>
ROM:0801317C                                         ; DATA XREF: ROM:08012C4Co
ROM:0801317C                                         ; 0
ROM:08013188                 LevelBackground <unk_851A368, unk_851E23C, unk_8523568> ; 1
ROM:08013194                 LevelBackground <unk_851A368, unk_851E23C, unk_8528AE8> ; 2
ROM:080131A0                 LevelBackground <unk_851A368, unk_851E23C, unk_852D57C> ; 3
ROM:080131AC                 LevelBackground <unk_851A368, unk_851E23C, unk_8531ED4> ; 4
ROM:080131B8                 LevelBackground <unk_851A368, unk_851E23C, unk_853614C> ; 5
ROM:080131C4                 LevelBackground <unk_851A368, unk_851E23C, unk_853A4C8> ; 6
ROM:080131D0                 LevelBackground <unk_851A368, unk_851E23C, unk_853E484> ; 7
ROM:080131DC                 LevelBackground <unk_851A368, unk_851E23C, unk_8542204> ; 8
ROM:080131E8                 LevelBackground <unk_851A368, unk_851E23C, unk_8545C3C> ; 9
ROM:080131F4                 LevelBackground <unk_851A368, unk_851E23C, unk_8549A60> ; $A
ROM:08013200                 LevelBackground <unk_851A368, unk_851E23C, unk_854DF24> ; $B
ROM:0801320C                 LevelBackground <unk_851A368, unk_851E23C, unk_8551D84> ; $C
ROM:08013218                 LevelBackground <unk_851A368, unk_851E23C, unk_855564C> ; $D
ROM:08013224                 LevelBackground <unk_851A368, unk_851E23C, unk_8559694> ; $E
ROM:08013230                 LevelBackground <unk_851A368, unk_851E23C, unk_855CF88> ; $F
ROM:0801323C LB_F0           LevelBackground <unk_8561434, unk_8561650, unk_8561774>
ROM:0801323C                                         ; DATA XREF: ROM:off_8012C50o
ROM:0801323C                                         ; 0
ROM:08013248                 LevelBackground <unk_8561C78, unk_8561E84, unk_8561FA8> ; 1
ROM:08013254                 LevelBackground <unk_8562478, unk_8562674, unk_8562798> ; 2
ROM:08013260                 LevelBackground <unk_8562C6C, unk_8562EE4, unk_8563008> ; 3
ROM:0801326C                 LevelBackground <unk_8563500, unk_85637A8, unk_85638CC> ; 4
ROM:08013278                 LevelBackground <unk_8563E24, unk_8563FEC, unk_8564110> ; 5
ROM:08013284                 LevelBackground <unk_8564654, unk_85648D0, unk_85649F4> ; 6
ROM:08013290                 LevelBackground <unk_8564EF4, unk_8565124, unk_8565248> ; 7
ROM:0801329C                 LevelBackground <unk_8565750, unk_85659AC, unk_8565AD0> ; 8
ROM:080132A8                 LevelBackground <unk_8565FD8, unk_8566260, unk_8566384> ; 9
ROM:080132B4                 LevelBackground <unk_8566898, unk_8566A7C, unk_8566BA0> ; $A
ROM:080132C0                 LevelBackground <unk_85670B0, unk_8567410, unk_8567534> ; $B
ROM:080132CC                 LevelBackground <unk_8567AA4, unk_8567CD4, unk_8567DF8> ; $C
ROM:080132D8                 LevelBackground <unk_8568300, unk_85684C0, unk_85685E4> ; $D
ROM:080132E4                 LevelBackground <unk_8568AF0, unk_8568CB4, unk_8568DD8> ; $E
ROM:080132F0                 LevelBackground <unk_85692E8, unk_856939C, unk_85694C0> ; $F
ROM:080132FC                 LevelBackground <unk_856999C, unk_8569A40, unk_8569B64> ; $10
END
