## 墨水屏实时天气服务端

参考 : [chenxuuu/luatos-eink-server: 供luatos墨水屏桌面摆件使用的服务端 ](https://github.com/chenxuuu/luatos-eink-server)

### HTTP参数

请求方法：GET

URL：`http://IP:22726/eink-weather/v1?l=112.4365,34.6069&t=XXXXXXXXXX&txt=跑路&date=20221231`

#### Query参数及说明

| 参数名 | 示例值           | 参数类型 | 是否必填 | 参数描述     |
| :----- | :--------------- | :------- | :------- | :----------- |
| l      | 112.4365,34.6069 | Text     | 是       | `位置`       |
| t      | XXXXXXXXXX       | Text     | 是       | `Token`      |
| txt    | 跑路             | Text     | 否       | `倒数日文本` |
| date   | 20221231         | Text     | 否       | `倒数日时间` |



### 天气API

网址：[彩云 | 登入 (caiyunapp.com)](https://dashboard.caiyunapp.com/user/sign_in/)  需要自己申请一个token



### 取坐标网站

取坐标网站：https://caiyunapp.com/map/?lang=cn?lang=cn#112.4365,34.6069

点击地图上的点，URL`#`后面的就是latlng坐标



### Lua程序

https://gitee.com/openLuat/LuatOS/blob/master/script/turnkey/eink-calendar/main.lua

第33行改为:

```lua
    local httpc = esphttp.init(esphttp.GET, "http://192.168.77.77:22726/eink-weather/v1?l=112.4365,34.6069&t=XXXXXXXXXX")
```

IP、token改为自己的

![image-20220727115015432](.\README.assets\image-20220727115015432.png)