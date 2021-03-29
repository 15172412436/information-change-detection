<template>
    <div>
        <body>
            <el-scrollbar style="height:100%;with:100%">
                <el-row>
                    <el-col :span="24">
                        <a class="order_title" style="padding: 0 10px;">首页</a>
                        <a class="order_title" style="padding: 0 10px;">版本更新</a>
                        <a class="order_title" style="padding: 0 10px;">关于</a>
                        <el-button round style="background-color: #2e86df;" @click="dialogFormLoginVisible = true">登录</el-button>
                        <el-button round @click="dialogFormRegisterVisible = true">注册</el-button>
                    </el-col>
                </el-row>
                <el-row>
                    <div style="width: 40%;height: 100%;float: left;">
                        <div class="grid-content">
                        <br>
                        <br>
                        <h1 style="text-shadow:2px 2px 1px wheat;font-family:PingFang SC;font-size: 80px">主题订阅</h1>
                        <p style="font-family: Hiragino Sans GB;font-size: 30px;">速度/快捷/方便</p>
                            <div  class="div_style" >
                                    订阅你的个人主题
                            </div>
                        </div>
                    </div>
                    <div style="position: absolutely;right:0px;width: 60%;height:100%;float:right">
                        <div style="margin:auto auto;height:100%;width:80%">
                            <el-carousel trigger="click" height="400px">
                                <el-carousel-item>
                                    <img src="../assets/images/bgleft.jpg" style="width: 100%; height: 100%">
                                </el-carousel-item>
                                <el-carousel-item>
                                    <img src="../assets/images/1.jpg" style="width: 100%; height: 100%">
                                </el-carousel-item>
                            </el-carousel>
                        </div>
                    </div>
                </el-row>
                <el-row>
                    <el-col :span="24">
                        <div style="text-align:center">
                        联系站长：xxx.xxx.xx;
                        </div>
                    </el-col>
                </el-row>
            </el-scrollbar>
                <!-- 以下就是弹出的登录框 -->
            <el-dialog style="width: 800px;margin: auto;" title="登录" :visible.sync="dialogFormLoginVisible" center>
                <el-form ref="Login_form" :model="Login_form">
                    <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
                    <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
                    <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
                    </el-form-item> -->
                    <el-tabs v-model="activeName_Login" >
                        <el-tab-pane label="账号密码登录" name="first">
                        <el-form-item label="用户名:" :label-width="formLabelWidth">
                            <el-input v-model="Login_form.name_log" autocomplete="off"></el-input>
                        </el-form-item>
                        <el-form-item label="密码:" :label-width="formLabelWidth">
                            <el-input v-model="Login_form.passwd_log" placeholder="请输入密码" show-password autocomplete="off"></el-input>
                        </el-form-item>
                        </el-tab-pane>
                        <el-tab-pane label="手机号快速登录" name="second">
                            <el-form-item label="手机号:" :label-width="formLabelWidth">
                                <el-input v-model="Login_form.phone_log" autocomplete="off" v-on:input="check_num('login')" ></el-input>
                            </el-form-item>
                            <el-form-item label="验证码:" :label-width="formLabelWidth">
                                <el-input v-model="Login_form.checkcode_log" autocomplete="off" style="width:120px" ></el-input>
                                <el-button type="primary" @click="send_code('login')" :disabled="is_disabled_Login"   style="font:5px">{{send_code_Login}}</el-button>
                            </el-form-item>
                        </el-tab-pane>
                    </el-tabs>
                </el-form>
                <span slot="footer" class="dialog-footer">
                    <el-button type="primary" @click="Login(activeName_Login)" >确 定</el-button>
                    <el-button @click="dialogFormLoginVisible = false" >取 消</el-button>
                </span>
            </el-dialog>

            <!-- 以下就是弹出的注册框 -->
            <el-dialog style="width: 800px;margin: auto;" title="注册" :visible.sync="dialogFormRegisterVisible" center>
                <el-form ref="Register_form" :model="Register_form">
                    <!-- <el-form-item label="验证码:" :label-width="formLabelWidth">
                    <el-input v-model="form.checkcode" style="width: 120px;" autocomplete="off"></el-input>
                    <img style="height: 40px; width: 80px;" src="./img/1.jpg"></img>
                    </el-form-item> -->
                    <el-tabs v-model="activeName_Register">
                        <el-tab-pane label="账号密码注册" name="first">
                            <el-form-item label="用户名:" :label-width="formLabelWidth">
                                <el-input v-model="Register_form.name_re" autocomplete="off"></el-input>
                            </el-form-item>
                            <el-form-item label="密码:" :label-width="formLabelWidth">
                                <el-input v-model="Register_form.passwd_re" placeholder="请输入密码" show-password autocomplete="off"></el-input>
                            </el-form-item>
                        </el-tab-pane>
                        <el-tab-pane label="手机号快速注册" name="second">
                            <el-form-item label="手机号:" :label-width="formLabelWidth">
                                <el-input v-model="Register_form.phone_re" autocomplete="off" ></el-input>
                            </el-form-item>
                            <el-form-item label="验证码:" :label-width="formLabelWidth">
                                <el-input v-model="Register_form.checkcode_re" autocomplete="off" style="width:120px"></el-input>
                                <el-button type="primary" @click="send_code('register')" :disabled="is_disabled_Register" style="font:5px">{{send_code_Register}}</el-button>
                            </el-form-item></el-tab-pane>
                    </el-tabs>
                </el-form>
                <span slot="footer" class="dialog-footer">
                    <el-button type="primary" @click.prevent="Register()" >确 定</el-button>
                    <el-button @click="dialogFormRegisterVisible = false" >取 消</el-button>
                </span>
            </el-dialog>
        </body>
    </div>
</template>

<script>
/* eslint-disable */
export default{
  name:'Index',
  data: function() {
      return { 
          activeName_Login: 'first',
          activeName_Register: 'first',
          send_code_Login:"发送验证码",
          send_code_Register:"发送验证码",
          is_disabled_Login:true,
          is_send_log:false,
          is_send_re:false,
          is_disabled_Register:true,
          visible: false,
          dialogFormLoginVisible: false,
          dialogFormRegisterVisible:false,
          Login_form: {
                  name_log: '',
                  passwd_log: '',
                  checkcode_log:'',
                  phone_log:'',
                  date1: '',
                  date2: '',
                  delivery: false,
                  type: [],
                  resource: '',
                  desc: ''
                  },
          Register_form: {
                  name_re: '',
                  passwd_re: '',
                  phone_re:'',
                  checkcode_re:'',
                  date1: '',
                  date2: '',
                  delivery: false,
                  type: [],
                  resource: '',
                  desc: ''
          },
          formLabelWidth: '80px'
           }
    },
  methods:{
    //登录函数
    Login:function(){
        var axios = require('axios');
        if(this.activeName_Login=='first'){
            var username = this.Login_form.name_log;
            var passwd = this.Login_form.passwd_log;
            var me = this;
            var data = {
                'username':username,
                'password':passwd
            };
            if(username=='' || passwd==''){
                me.$message.error('登录失败:用户名或密码为空');
            }
            else{
            //请求参数
            var config = {
                method: 'post',
                // 用nginx进行代理配置 把所有外部链接全部改成此端口出发的链接
                // url: 'http://175.24.28.202:8000/auth/login',
                url: 'http://175.24.28.202:80/login',
                data : data,
                headers: { 
                    'Access-Control-Allow-Origin':'*'
                },
                };
            //发出请求
            axios(config)
            .then(function(response){
                //成功失败在这儿
                var return_data = response.data;
                if(return_data.status=='1'){
                    me.$message({
                    message: '登录成功',
                    type: 'success'
                    });
                    localStorage.setItem('token',return_data.data.token);
                    var userid = return_data.data.uid;
                    if(return_data.data.subs.dtdy==null)
                    {
                        var dtdy_s = [];
                    }
                    else{
                        var dtdy_s = return_data.data.subs.dtdy;
                    }
                    if(return_data.data.subs.fcdy==null){
                        var fcdy_s = [];
                    }
                    else{
                        var fcdy_s = return_data.data.subs.fcdy;
                    }
                    if(return_data.data.subs.ztdy==null){
                        var ztdy_s = [];
                    }
                    else{
                        var ztdy_s = return_data.data.subs.ztdy;
                    }

                    me.$router.push({path:'/main',query:{username:userid}});
                }
                else{
                    me.$message.error('登录失败:'+return_data.msg);
                }
            })
            .catch(function (error) {
                //出了问题就在这儿
                console.log(error);  
            });
            }
        }
        else{
            var phone = this.Login_form.phone_log;
            var code = this.Login_form.checkcode_log;
            var me = this;
            var data = {
                'phone': phone,
                'code':code
                };
            if(phone=='' || code == ''){
                me.$message.error('登录失败:手机号或验证码为空');
            }
            else{
            //请求参数
            var config = {
                method: 'post',
                // http://175.24.28.202:8000/auth/phoneLogin;
                url: 'http://175.24.28.202:80/phoneLogin',
                data : data,
                headers: { 
                    'Access-Control-Allow-Origin':'*'
                },
                };
            //发出请求
            axios(config)
            .then(function(response){
                //成功失败在这儿
                var return_data = response.data;
                console.log(return_data);
                if(return_data.status=='1'){
                    me.$message({
                    message: '登录成功',
                    type: 'success'
                    });
                    localStorage.setItem('token',return_data.data.token);
                    me.$router.push({path:'/main',query:{username:return_data.data.uid}});

                }
                else{
                    me.$message.error('登录失败:'+return_data.msg);
                }
                
                
            })
            .catch(function (error) {
                //出了问题就在这儿
                console.log(error);
                
                
            });
            }
            
        }
        //this.dialogFormLoginVisible = false;
        //this.$router.push({path:'/main'}); 
    },
    //注册函数
    Register:function(){
        var axios = require('axios');
        if(this.activeName_Register=='first'){
            var username = this.Register_form.name_re;
            var passwd = this.Register_form.passwd_re;
            var me = this;
            var data = {
                "username":username,
                "password":passwd
            };
            if(username==''||passwd==''){
                me.$message.error('注册失败:用户名或密码为空');
            }
            else{
            //请求参数
            var config = {
                method: 'post',
                // http://175.24.28.202:8000/auth/register;
                url: 'http://175.24.28.202:80/register',
                headers: { 
                    'Access-Control-Allow-Origin':'*'
                },
                data : data
                };
            //发出请求
            axios(config)
            .then(function(response){
                //成功失败在这儿
                var return_data = response.data;
                if(return_data.status=='1'){
                    console.log(return_data);
                    me.$message({
                    message: '注册成功',
                    type: 'success'
                    });
                    me.$router.push('/index');
                }
                else{
                    me.$message.error('注册失败:'+return_data.msg);
                }
            })
            .catch(function (error) {
                //出了问题就在这儿
                console.log(error);  
            });
            }
        }
        else if(this.activeName_Register=='second'){
            var phone = this.Register_form.phone_log;
            var code = this.Register_form.checkcode_log;
            var me = this;
            var data = {
                'phone': phone,
                'code':code
                };
            if(phone=='' || code == ''){
                me.$message.error('登录失败:手机号或验证码为空');
            }
            else{
                //请求参数
                var config = {
                    method: 'post',
                    // http://175.24.28.202:8000/auth/phoneLogin;
                    url: 'http://175.24.28.202:80/phoneLogin',
                    data : data,
                    headers: { 
                    'Access-Control-Allow-Origin':'*'
                    },
                    };
                //发出请求
                axios(config)
                .then(function(response){
                    //成功失败在这儿
                    var return_data = response.data;
                    console.log(return_data);
                    if(return_data.status=='1'){
                        console.log(return_data);
                        me.$message({
                        message: '注册成功',
                        type: 'success'
                        });
                        me.$router.push({path:'/index'});
                    }
                    else{
                        me.$message.error('注册失败:'+return_data.msg);
                    }
                    
                    
                })
                .catch(function (error) {
                    //出了问题就在这儿
                    console.log(error);
                    
                    
                });
            }
        }
      
      this.dialogFormRegisterVisible = false;
      
    },
    //发送验证码
    send_code:function(type){
        const qs = require('qs');
        const axios = require('axios');
        if(type=='login'){
            var num = 60;
            this.is_disabled_Login = true;
            this.is_send_log = true;
            var buttonName = this.send_code_Login;
            var me = this;
            let interval = setInterval(function() {
                //alert(num);
                me.send_code_Login =num+ "秒后重新发送";
                --num;
                if(num < 0) {
                    me.send_code_Login = "重新发送";
                    me.is_disabled_Login = false;
                    me.is_send_log=false;
                    clearInterval(interval);
                }
            }, 1000);
            var phone_num = this.Login_form.phone_log;
            var data =  qs.stringify({
                    'phone': phone_num
                    });
            //请求参数
            var config = {
                method: 'post',
                // http://175.24.28.202:8000/auth/code;
                url: 'http://175.24.28.202:80/code',
                headers: { 
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                data : data
                };
            //发出请求
            axios(config)
            .then(function(response){
                //成功失败在这儿
                console.log(response);
                var return_data = response.data;
                if(return_data.status=='1'){
                    me.$message({
                        message: '验证码发送成功',
                        type: 'success'
                        });
                }
                else{
                    me.$message.error('验证码发送失败:'+return_data.msg);
                }
                console.log(return_data.code);
                
            })
            .catch(function (error) {
                //出了问题就在这儿
                console.log(error);  
            });
        }
        else if(type=='register'){
            var num = 60;
            this.is_disabled_Register = true;
            this.is_send_re = true;
            var me = this;
            let interval = setInterval(function() {
                //alert(num);
                me.send_code_Login =num+ "秒后重新发送";
                --num;
                if(num < 0) {
                    me.send_code_Register = "重新发送";
                    me.is_disabled_Register = false;
                    me.is_send_re = false;
                    clearInterval(interval);
                }
            }, 1000);
            var phone_num = this.Register_form.phone_re;
            var data =  qs.stringify({
                    'phone': phone_num
                    });
            //请求参数
            var config = {
                method: 'post',
                // http://175.24.28.202:8000/auth/code;
                url: 'http://175.24.28.202:80/code',
                headers: { 
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'Access-Control-Allow-Origin':'*'
                },
                data : data
                };
            //发出请求
            axios(config)
            .then(function(response){
                //成功失败在这儿
                console.log(response);
                var return_data = response.data;
                if(return_data.status=='1'){
                    me.$message({
                        message: '验证码发送成功',
                        type: 'success'
                        });
                }
                else{
                    me.$message.error('验证码发送失败:'+return_data.msg);
                }
                console.log(return_data.code);
                
            })
            .catch(function (error) {
                //出了问题就在这儿
                console.log(error);  
            });
        }
        
    },
    //onkeyup判断手机号码
    check_num:function(type){
        if(type=='login'){
            var val = this.Login_form.phone_log;
            var r = /((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}/g;
            //alert(this.Login_form.phone_log.replace('^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$'));
            if(val.length>11){
                this.is_disabled_Login=true;
            }
            else{
                if(r.test(val))
                {   
                    if(this.is_send_log==true){
                        this.is_disabled_Login=true;
                    }
                    else{
                        this.is_disabled_Login=false;
                    }
                    
                }
                else{
                    this.is_disabled_Login=true;
                }
            }
        }
        else if(type=='register'){
            var val = this.Register_form.phone_re;
            var r = /((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}/g;
            //alert(this.Login_form.phone_log.replace('^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$'));
            if(val.length>11){
                this.is_disabled_Register=true;
            }
            else{
                if(r.test(val))
                {   
                    if(this.is_send_re==true){
                        this.is_disabled_Register=true;
                    }
                    else{
                        this.is_disabled_Register=false;
                    }
                    
                }
                else{
                    this.is_disabled_Register=true;
                }
            }
        }
    }
  }
}
</script>
<style>
.el-col {
height: 40px;
}
.bg-purple {
height: 100%;
background: #d3dce6;
}
.bg-purple-light {
background: #e5e9f2;
}
.grid-content {
border-radius: 4px;
min-height: 600px;
height: 90%;

text-align: center;
margin: auto auto;

}
.row-bg {
padding: 10px 0;
background-color: #f9fafc;
}
body{
height: 100%;
width: 100%;
background: url('../assets/images/indexbg.jpg') no-repeat;
background-size: cover;
position: absolute;
overflow: hidden;
}
.bgima-left{

height: 100%;
width: 100%;
background: url('../assets/images/bgleft.jpg') no-repeat;
background-size: cover;
position: absolute;
overflow: hidden;

}
.bg-orange{
background: orange;
border-radius: 4px;
text-align: center;

}
.div_style{
align-self: center;
font-family: Hiragino Sans GB;
font-size: 30px;
color: white;
border: 1px solid black;
border-radius: 20px 20px 20px 20px;
background: orange;
border-width: 1px;
width: 300px;
margin: 0 auto;
}
.order_title{
font-family: Hiragino Sans GB;
font-size: 20px;
color: black;
}
</style>