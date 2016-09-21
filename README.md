# beego_api
一分钟创建应用api


1.在本地数据库创建一个数据库go,建一个表
    
    DROP TABLE IF EXISTS `user`;
    CREATE TABLE `user` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `username` varchar(20) DEFAULT NULL,
      `password` varchar(20) DEFAULT NULL,
      `age` tinyint(2) DEFAULT NULL,
      `email` varchar(30) DEFAULT NULL,
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
    
    插入两条数据
    INSERT INTO `user` VALUES ('1', 'baijun', '123456', '20', '6666@qq.com');
    INSERT INTO `user` VALUES ('2', 'baijun66', '12345699', '26', '888@qq.com');
    
  2.在src/下bee api beego_api -conn=root:@tcp(127.0.0.1:3306)/go,创建api项目，已生成user表相关（controller,model,router规则），然后cd到beego_api目录，执行bee run -gendoc=true -downdoc=true开启应用，http://127.0.0.1:8888/swagger/  查看即可测试api应用 user的crud操作
