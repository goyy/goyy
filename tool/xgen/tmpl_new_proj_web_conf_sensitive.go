// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebConfSensitive = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<sensitive name="<%.NewProjName%>">
				<enable>false</enable>
				<excludes></excludes>
				<values></values>
			</sensitive>
		</environment>
		<environment id="test">
			<sensitive name="<%.NewProjName%>">
				<enable>false</enable>
				<excludes></excludes>
				<values></values>
			</sensitive>
		</environment>
		<environment id="production">
			<sensitive name="<%.NewProjName%>">
				<enable>true</enable>
				<excludes>
				<![CDATA[
				/,
				/login
				]]>
				</excludes>
				<values>
				<![CDATA[
				getWriter,FileOutputStream,getRuntime,getRequest,getProperty,script,frameset,iframe,<marquee,<object,document.,.cookie,.href,alert(,confirm(,prompt(,expression(,$.get,$.post,$.ajax,touchstart,touchmove,touchend,touchcancel,gesturestart,gesturechange,gestureend,onorientationchange,orientationchange,
				onabort,onafterprint,onbeforeonload,onbeforeprint,onbeforeunload,onblur,oncanplay,oncanplaythrough,onchange,onclick,onconte,oncontextmenu,ondblclick,ondrag,ondragend,ondragenter,ondragleave,ondragover,ondragstart,ondrop,ondurationchange,onemptied,onended,onerror,onfocus,onformchange,onforminput,onhaschange,oninvalid,oninput,onkeydown,onkeypress,onkeyup,onload,onloadeddata,onloadedmetadata,
				onloadstart,onmessage,onmousedown,onmousemove,onmouseout,onmouseover,onmouseup,onmousewheel,onoffline,ononline,onpagehide,onpageshow,onpause,onplay,onplaying,onpopstate,onpropertychange,onprogress,onratechange,onreadystatechange,onredo,onreset,onresize,onscroll,onseeked,onseeking,onselect,onstalled,onstart,onstorage,onsubmit,onsuspend,ontimeupdate,onundo,onunload,onvolumechange,onwaiting,
				<%if iszhcn%>援交,色狼集中营,狼客娱乐网,AV种子,肉漫,上门服务,上门援交,哪里有小姐,保健品数据,电视购物数据,快递面单数据,快递面单资料,车主资源,业主数据,车主数据,电购数据,业主资源,收藏品数据,业主资料,车主资料,老人数据,股民数据,股民资源,报仇公司,PCP枪管膛线,秃鹰扳机,铅弹母鸡,华尔街娱乐城,天成国际娱乐城,天猫娱乐城,大中华娱乐城,嘉年华娱乐城,米其林娱乐城,喜力国际娱乐城,豪门国际俱乐部,凯旋门娱乐城,永辉国际娱乐城,太阳城百家乐,万达国际娱乐城,鼎丰国际娱乐城,帝王娱乐城,宝龙娱乐城,皇冠娱乐城,
				天上人间,黄赌毒,金牌娱乐城,半岛国际娱乐城,澳博国际娱乐城,五湖四海娱乐城,澳门星际赌场,威龙国际娱乐城,帝豪国际娱乐城,鼎龙国际娱乐城,太阳城,百家乐平台,仿真硅胶面具,液体哌甲酯,甲氧麻黄酮,甲基甲卡西酮,大马士革刀,军刀,廓尔喀军刀,勃朗宁军刀,格斗刀,伞兵刀,戈博刀,战术匕首,三棱军刀,战术军刀,狗腿弯刀,三棱刺刀,三棱尖刀,弹簧刀,匕首,三棱刮刀,卡巴军刀,双刃尖刀,虎牙刀,尼泊尔弯刀,跳刀,三棱尖刺,兰博刀,武士刀,战术折刀,身份证制作软件,身份证复印件生成软件,身份证号码生成软件,身份证复印件生成器,曾道人特码,身份证号码生成器,乐透,轮盘,德州扑克,
				扎金花,梭哈,棋牌,赌球,赌马,特码,斗地主,娱乐城,百家乐,彩金,博彩,六合彩,代办,克隆,办证,投注,下注,开奖,身份证模板,3d轮盘,八肖中特,盘口高额返水,女宝,女童,男童,幼童,男婴,男宝,专业办证公司,刻章办证公司,文凭代办网,证件制作网,办证刻章公司,办证公司网,办证工作室,办证服务公司,证件网,办证刻章,英语替考,办证网,四六级替考,替考网站,四六级助考,枪手网站,枪手替考,GFX面具,飞叶子,纯缅麻古,黄冰,朝鲜冰,钻石冰,帮人报仇,专业报仇,替人了仇,职业报仇,帮人出气,
				找职业杀手,替人消灾,替人追债,替人出气,雇人复仇,买凶杀人,替人复仇,替人报仇,提供打手,雇打手网,替人报仇网,找人报仇网,帮人复仇网,特洛伊卧底软件官方网,专业删帖机构,x手机卧底软件官网,手机监听官网,网络删贴公司,三利达官网,弩弓官网,军刀网,弓弩网,户外军品网,代办发票公司,代理发票公司,发嘌,发缥,发剽,莫洛托夫鸡尾酒,红烧兔子大餐,发漂,无政府主义者食谱,3d打印枪支图纸,改火套件,小六改火,PCP气枪网,枪模网,猎枪销售网,仿真枪械网,气枪专卖网,进口气枪网,气狗专卖网,工字气枪网,三箭气枪网,猎枪网,套牌车,气枪网,车辆牌照,车牌套牌,国安证,
				言正,军官证 ,警官证 ,氰化钠,溴化铯,山奈,氰化银,氰化钾,氰化钙,乙醚,硝酸汞,砷化氢 ,亚砷酸钾,重铬酸钠,氧化铊,乙酸亚铊,丙二酸铊,四氧化锇,乙硼烷,六氟丙酮,三氟化硼,番木鳖碱,氯磺酸,马钱子碱,碳酸亚铊,丙腈,甲基肼,丁腈,异丁腈,烯丙胺,亚硝酸乙酯,氯甲酸甲酯,乌头碱,一氯乙醛,氯乙酸,丙烯醛,乙酸苯汞,放线菌酮,二盐酸盐,地高辛,五氯酚钠,甲藻毒素,赭曲毒素,二氯化汞,硫酸亚铊,氯化汞,乙酸汞,溴化汞,羰基氟,丙二酸鉈,银氰化钾,碘甲烷,碘化汞,三氯化砷,氰化物,
				氯仿,白砒,醋酸铊,砷酸氢汞,蓖麻毒素,氧化汞,铊盐,砒霜 ,山奈钾,氰化汞,哌甲酯,硝酸铊,苯巴比妥,金属铊,氯化锰,当面交易,氰化锌,碘化氰,套号,公务员考试答案,考研答案,国考答案,高考考中答案,高考答案,考试答案,中考考中答案,高考考前答案,中考考前答案,无线电作弊器材,考前答案,反屏蔽考试设备,四六级答案,考试作弊设备,考试作弊器材,考中答案,英语等级考试答案,考试作弊工具,针孔作弊器,警用臂章,仿真警服,考试题,考试作弊器,警帽,武警作战服,警用器材,特警作战服,手铐,警用甩棍,警衔,警徽,警用钢叉,警服,警灯,警察作训服,
				交警警服,警察胸标,警察执勤服,电警棍,高压电警棍,高仿警服,宅急送数据,警察证,银行客户名录,婴儿信息,银行卡用户资料,银行卡用户信息,业主身份资料,业主信息,业主名单,业主身份信息,学生家长资料,学生档案,学生家长名单,学生家长名录,信用卡客户资料,物流客户数据,网购客户资料,物流客户资料,顺丰快递数据,顺丰面单数据,全球通用户资料,收藏品客户资料,社保资料,期货客户名单,期货客户资源,期货客户资料,女性数据,女性资料,拍拍用户资料,落榜考生名单,楼盘业主数据,楼盘业主资料,老年人资料,楼盘业主名单,联通客户资料,老板通讯录,老年人数据,老年人信息,
				老板手机号码,客户名单,客户信息,客户数据,家长数据,考生资源,金融客户资源,户主资料,户主信息,家长资料,股民信息,股民资料,股民名单,股民联系方式,股民名录,股民开户数据,股民个人资料,股民个人信息,高消费人群名录,股民电话资源,股民电话号码,高考学生信息,富人资料,高官名录,富豪数据,富人信息,富人数据,服刑人员资料,房地产客户资料,犯人数据,房主数据,法人资料,电视购物资源,电信用户资料,法人数据 ,电视购物名录,电购面单数据,别墅业主信息,车主名录,车主信息,保险客户名录,保险客户资料,毕业生简历,保险客户名单,
				保健品客户资料,高仿真脸皮,硅胶面具,乳胶人皮面具,易容面具,高仿真人皮面具,无声手枪,微声手枪,冲锋枪,微冲,微型冲锋枪,瓦斯枪,双筒猎枪,热兵器,制式手枪,改装发令枪,改装射钉枪,枪械,来复枪,枪模,电击枪,仿真气枪,打鸟气枪,运动步枪,打鸟枪,打鸟汽枪,双管猎枪,汽步枪,短枪,平式枪,真枪,枪,长枪,狙击枪,火药枪,火枪,警用枪支,麻醉枪,军用手枪,猎枪,仿真枪,气步枪,气枪,快删,手枪,步枪,高仿钞,硬币模具,假币模具,伪币,伪钞,假钱,钞票,氢弹,巡航导弹,迫击炮弹,
				硅烷炸弹,汽油弹,手雷,土炸弹,导弹,催泪弹,黑索今,液体炸弹,烟雾弹,黄色炸药,太恩,硝胺炸药,硝酸胺炸药,手榴弹,硝铵炸药,太安,水胶炸药,三硝基苯酚,乳化炸药,三硝基甲苯,内裤炸弹,雷汞,銷售,黑索金,导爆索,铵油炸药,铵梯炸药,季戊四醇四硝酸酯,硝化甘油,原子弹,TNT,PETN,塑料炸弹,燃烧瓶,燃烧弹,雷管,黑火药,炸药,炸弹,xwodi软件,手机监听王,移动电话卧底软件,卧底定位软件,智能偷听软件,x卧底软件,spyera软件,特洛伊卧底软件,手机远程监控软件,手机窃听软件,手机监听软件,手机间谍软件,手机卧底软件,信息快照,卧底监控软件,负面贴文,论坛信息,
				负面评价,造谣信息,诽谤信息,不利信息,陪睡,天涯帖,负面视频,负面消息,负面新闻,百度贴吧帖子,负面论坛贴,负面帖子,负面微博,负面信息,负面评论,磁道信息,银行卡号信息,磁条信息,信用卡磁条,银行卡磁条,借记卡磁条,黑卡,银行卡解码器,磁道内码,银行卡复制器,折叠手弓弩,森林之狼弓弩,弩箭,战神弓弩,反恐弩,力斯曼弩,二手弩,小飞狼弩,秦氏弓弩,有偿捐赠,有偿捐献,弓弩,弩弓,进口弓弩,追风弓弩,猎豹弓弩,眼镜蛇弩,有偿提供,大黑鹰弓弩,赵氏弓弩,三利达弓弩,黑曼巴弩,阻击弩,军用弓弩,军用折叠弩,狩猎弩,军用弩箭,警用弩,手弩,
				钢弩,踏弩,军用钢珠弩,十字弩,弩枪,56式枪刺,56式刺刀,大马士革军刀,军用刺刀,军刺枪刺,野战砍刀,军用潜水刀,七孔狗腿刀,战术军刺,武士直刀,战术刀具,d80军刺,丛林开山刀,军品刺刀,战术狗腿刀,阿普唑仑片,K粉,ghb水,氨茴酸,安眠酮,安非他明,苯乙胺,白粉,安钠咖,苄基丙酮,吡利啶,苯异丙胺,冰毒,丙吡兰,地西泮,德美罗,大麻,二甲基安非他明,古柯碱,呋替啶,氟硝西泮,海乐神,胡椒基甲基,胡椒醛,海洛因,酣乐欣,黄麻素,黄体酮,甲基可可碱,甲基麻黄素,甲基安非他明,甲基苯丙胺,咖啡因,甲卡西酮,甲硝西泮,可待因,可卡因,卡西酮,黎城辣面,利多尔,六氢大麻酚,氯氨酮,力月西片,
				麻古,氯胺酮,氯硝西泮,麻黄素,麻果,麻黄碱,麻黄浸膏,麻黄素羟亚胺,麦角醇,麦司卡林,麦角酸,麦角乙二胺,莫达非尼,尼美西泮,尼二氢可待因,尼可待因,尼蒙尔克素,普拉西泮,普斯普剂,青蒿素 ,氢可酮,去甲麻黄素,去氧麻黄碱,去甲伪麻黄碱,三唑仑,沙菲片,双氢可待因,双氢吗啡,天然咖啡因,替马西泮,香港ghb,伪麻黄素,香港GHB粉,盐酸氯,盐酸丁丙诺啡,亚甲基二氧苯基,盐酸氯胺酮,盐酸麻黄碱,盐酸麻黄素,盐酸曲马多,盐酸哌替啶,盐酸羟亚胺,应苄基丙酮,已环利定,乙基吗啡,左啡诺,长治筋,左美沙芬,成人3d电视棒,AV电视棒,成人电视棒,高清成人电视棒,成人AV电视棒,3D网络电视棒成人版,
				高清3d成人电视棒,眼角膜,成人3d网络电视棒,shen源,肾器官,肾脏,肝源,肝脏,肾源,现货,求购,有偿卖,有偿献,有偿售,援交妹,甲氰菊酯,马拉硫磷,羟基乙腈,慢性毒药,安乐死毒药,士的宁,蓖麻毒素,三氯化二砷,氰化铊,绿藜芦生物碱,琥珀胆碱,敌百虫,氮芥气,致癌慢性毒药,立克命,百治磷,二甲基硫代磷酰氯,甲基丙烯腈,二乙基硫代磷酰氯,甲基肼,丙腈,丁腈,异丁腈,无水肼,三氯甲烷,西力生,乙基氰,溴氰菊酯,四氯硫代碳酰,毒鼠强,氯硫酸,氯甲酸甲酯,硝酸铊,甲基磺酰氟,氯气,全氟丙酮,三氯硫氯甲烷,
				光气,全氟异丁烯,沙林,氯化硫酰甲烷,过氯甲硫醇,氯化硫酸,液氯,索曼,氟醋酸钠,甲磺酰氟,塔崩,毒药,氰化钾,氰化金钾,三氯化砷,甲烷磺酰氯,LSD迷幻药,安非他明,左旋麻黃素,胡椒醛,氯胺酮,GHB水,麻黃素,胡椒基甲酮,甲卡西酮,左旋麻黄素,可待因,鸦片,甲基安非他命,酣乐欣,摇头丸配方,盐酸羟亚胺 ,安眠酮,盐酸麻黄碱,海乐神,天然咖啡因,天然可卡因,去氧麻黄素,麦角胺,麦角酸二乙胺,麻古制作方法,氟硝西泮,k粉制作技术,甲硝西泮,k粉制作方法,氯胺酮制作配方,中华打猎论坛,新千秋狩猎论坛,诚信汽枪网,中华狩猎论坛,中国汽枪网,52名刀网,高压汽枪,
				真钱轮盘赌博,蓝盾国际娱乐城,蒙特卡罗娱乐城,新葡京娱乐城,盈丰国际娱乐城,黄金城娱乐场,现金棋牌游戏,乐九娱乐城,bet365娱乐场,战神国际娱乐城,圣淘沙娱乐城,大赢家真人娱乐,百家乐,轮盘,赌球,最新银行卡复制器,假身份证,军官证样本,假军官证,高仿身份证,假文凭,假户口本,高仿户口本,假结婚证,高仿学位证,高仿人民币,弩枪,森林之虎,钢弩,十字弓弩,户撒刀,蝴蝶刀,阿昌刀,邻氨基苯甲酸,甲苯喹唑酮,亚甲二氧基甲基苯丙胺,甲基苯丙胺,相思红娘粉,自制K粉教程,云南情蛊,麻果,麻果配方,盐酸氯胺酮,代删,身份证复印件制作器,身份证复印件模板,红9娱乐城,91娱乐城,必胜国际,
				老钱庄娱乐,皇朝娱乐城,金海岸娱乐城,一码中特,皇冠平台,六合玄机,冰毒的制作方法,毕业文凭,红樱桃催情液,汽车电子干扰器,假警官证,高仿军官证,秃鹰汽枪网,汽车电子解码器,麻古,金盛国际娱乐城,真钱百家乐,永利高娱乐城,百乐坊娱乐城,在线现金扑克,菲律宾太阳城,凯撒皇宫娱乐城,百丽宫娱乐城,金沙娱乐场,金木棉娱乐城,嘉禾娱乐城,永利娱乐城,美高梅娱乐城,金沙娱乐城,E世博娱乐城,身份证复印件生成器,身份证号码生成器,假币,假钞,森林之狼,小飞狼,森林之鹰,蜘蛛刀,三棱刀,博伊刀,廓尔喀刀,苗刀,直刀,卡巴刀,蝴蝶甩刀,人造芥子油,唛可奈因,苯环利定,迷幻蘑菇,美沙酮,恩华三唑仑,
				盐酸羟亚胺,硝甲西泮,植物冰,麻黄碱,黄牙签,k粉,二代身份证制作软件,假证件,学历证书,套号毕业证,套牌车,国外文凭,黑火药,土制手雷,土制炸弹,土制手榴弹,女性资料,女性数据,男性数据,进口气狗,渔夫吹图纸,快排吹教程,diy吹尘,女童,狗腿刀,丛林刀,巴克刀,丁丙诺啡,可卡因,盐酸哌替啶,海洛因,摇头丸,氟硝安定,安纳咖,肉文,快排吹图纸,土铳,进口汽枪,国产气枪,钢珠枪,沙漠之鹰,老人资料,学生数据,学生资料,宅急送数据,删帖公司,二代身份证生成器,普通增值税发票,机打发票,军车套牌,假车牌,三利达,大黑鹰,麻醉弩箭,缅果,苯丙胺,倍它洛尔,麻黄素,福寿膏,沙菲片,三唑仑,
				大麻烟,GHB原液,香料精灵,安钠咖,亚甲二氧甲基苯丙胺,曲马多,朝鲜冰,开山刀,廓尔喀弯刀,弹簧跳刀,三棱跳刀,丛林直刀,勃朗宁刀,潜水刀,西布曲明,乙基吗啡,帮人复仇,报仇公司,找杀手公司,打手网,杀手网,打火机吹,吹尘diy,AV电视棒,成人版电视棒,成人高考考前答案,考试作弊器,硬币,伪钞,伪币,假rmb,jia币,锡纸万能开锁,快开工具,万能钥匙,万能开琐工具,万能开锁器,强开工具,万能开锁工具,开锁器,汽车信号解码器,汽车遥控解码器,汽车信号屏蔽器,偷听软件,x卧底软件,监听软件,汽车钥匙解码器,汽车解码器,汽车遥控拦截器,汽车拦截器,汽车干扰器,手机监听软件,
				手机007间谍软件,特洛伊卧底软件,卧底定位软件,监听器软件,xwodi软件,手机GSM监听设备,手机监听设备,警用电击棒,警用臂章,高压电警棍,高压警棍,警用甩棍,警察徽章,警用电棒,警用手铐,高仿警官证,金融客户资料,楼盘业主资料,家长资料,收藏品客户资料,保险客户资料,电视购物资料,保险客户数据,小区业主名录,股民开户数据,物流客户数据,老年人资料,美秃套件,汽狗配件,金钟汽枪,金钟气狗,五四式手狗,54式手木仓,64式气木仓,92式手木仓,64式手木仓,联盟狩猎论坛,东方狩猎论坛,三角牌气枪,汽狗狗粮,七七式手枪,工字气qiang,高压鸟枪,狗管,仿真汽枪,峨眉牌汽枪,老妖狩猎论坛,
				气枪枪管,气狗狗粮,上海工字气枪,钢珠猎狗,工字汽狗,锡锋牌气枪,骚本套件,猎枪铅弹,汽狗铅弹,气枪瞄准镜,气枪子弹,汽枪配件,枪支配件,五四式手枪,秃鹰套件,汽枪铅弹,钢珠弹,军用枪支,改装发令枪,改装射钉枪,德版PPK,温切斯特气枪,手拉狗,上海工字汽狗,1比1仿真枪,仿真气狗,三箭气狗,上海工字牌气狗,三箭汽枪,手枪,仿真气木仓,猎狗铅弹,三箭汽狗,锡锋汽枪,枪模,单管猎枪,汽枪消声器,雷鸣登猎枪,雷鸣登汽枪,猎枪配件,打鸟枪,三箭气枪,手狗配件,瓦斯枪,金属仿真枪,金属枪模,高压打鸟枪,气步枪,自制手枪,高压步枪,气狗配件,一比一枪模,秃鹰汽枪,健卫步枪,汽短狗,狙击狗,土制猎枪,
				台秃,汽长狗,左轮手枪,狙击枪,45MM狗粮,45MM铅弹,92式手枪,92式手狗,77式手狗,77式手枪,654k手枪,54式手枪,654K气枪,钢珠狗,CFX气枪,pcp汽枪,pcp气枪,小姐威客网,丝袜恋足会所,楼凤信息,兼职小姐,淫淫网,就去射,AV网址,性息平台,无码成人影院,国产A片下载,激情少妇,援交会所,午夜成人电影,亚洲色站,裸聊,性爱电影,淫图,寂寞少妇,清纯学生妹,菀式服务,丝袜美女,毛片网址,三级片地址,丐帮情色论坛,少妇白洁,枪刺,巴克折刀,砷酸汞,敌杀磷,敌恶磷,安乐死药物,氯化银钾,氯仿,黄曲霉素,氰化银,山埃,
				麻古配方,甲基苯丙胺制作方法,海洛因制作方法,冰毒技术,苯基甲酮,古可叶,咖啡碱,缅古制作方法,麦角酸提炼技术,缅古配方,缅古制作技术,脱氧麻黄碱,摇头丸制作方法,罂粟碱,咖啡因,古柯碱,朝鲜冰,甲基乙基酮,异黄樟素,黄樟油,冰毒提炼教程,求购,黎城辣面,身份证号生成器,猎弩,军用弩,手弩,三利达追月,三利达大黑鹰,麻醉箭弩,上海工字气狗,骚本气枪,54式图纸,手枪见面交易,枪械制造方法,汽枪专卖网,气枪网,天朝狩猎论坛,打鸟汽枪,钢珠汽枪,TNT炸药,手榴弹,导火索,定时炸弹,土炸药,电子炸弹,鱼雷,遥控炸弹,土制燃烧弹,黑索金,土制火药,C4橡胶炸药,自制炸药,爆破剂,自制炸弹,液体炸药,
				硝胺炸药,工业雷管,军用雷管,高爆炸药,任意改号软件,手机改号软件破解版,91电影网,怡红电影网,删帖,投注娱乐,大三元,洗码返水,中国教父习近平,习近平的政治公关,习近平和毛泽东,习李争霸战,习近平恶斗团派,习近平遇刺始末,习近平家族的海外秘密资产,周永康余党,周永康之后轮到谁,习近平王岐山围捕周永康家族,胡春华大战性都,汤灿引爆打老虎,中国情色游戏,中国大风险,方励之纪念文集,萤火虫的反抗——这个世纪的知识分子,平等团结路漫漫——对我国民族关系的反思,爆破工程技术人员挂靠,爆破安全作业工程师挂靠,爆破工程师挂靠,爆破证挂靠,爆破挂靠,
				九评共产党,退党保平安,<%end%>
				]]>
				</values>
			</sensitive>
		</environment>
	</environments>
</configuration>
`
