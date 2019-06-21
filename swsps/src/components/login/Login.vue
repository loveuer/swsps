<template>
    <div id="container">
        <div class="surface">
            <div class="left">
                <div @click="meetHome">
                    <div style="color:lightblue;">Safety</div>
                    <div style="color:lightgreen;">Wing</div>
                    <div>SPS</div>
                </div>
                
            </div>
            <div class="right">
                <div class="login-zone">
                    <div>
                        <div>
                            <el-input v-model="username">
                                <i slot="prefix" class="el-input__icon el-icon-user-solid"></i>
                            </el-input>
                        </div>
                        <div>
                            <el-input v-model="password" show-password>
                                <i slot="prefix" class="el-input__icon el-icon-lock"></i>
                            </el-input>
                        </div>
                        <div>
                            <el-row>
                                <el-button type="primary" @click="dologin" :loading="logining">登录</el-button>
                                <el-button @click="sorryfornotwork">注册</el-button>
                            </el-row>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: '',
            password: '',
            logining: false,
        };
    },
    methods: {
        meetHome() {
            this.$router.push('/');
        },
        sorryfornotwork: function() {
            this.$message({showClose: true, message: '已支持用事务管理的账户登录', type: 'info',});
        },
        dologin() {
            if (!this.username || !this.password) {
                this.password = '';
                this.$message({showClose: true, message: '用户名和密码不能为空', type: 'warning',});
                return;
            };
            this.logining = true;
            this.$http.post("/api/user/login", `username=${this.username}&password=${this.password}`)
                .then(resp => {
                    if (resp.data.id !== 0) {
                        this.$store.commit("setUser", resp.data.user);
                        this.setCookie("sessid", resp.data.sessid, 60);
                        this.$router.push("/");
                    } else {
                        this.password = '';
                        this.logining = false;
                        this.$message({showClose: true, message: '用户名和密码不匹配', type: 'warning',});
                        return;
                    }
                })
                .catch(err => {
                    this.password = '';
                    this.logining = false;
                    this.$message({showClose: true, message: '发生未知错误, 请刷新页面重试', type: 'warning',});
                    console.log(err.response);
                    return
                });
        },
        setCookie: function(cname, cvalue, expire_min) {
            var d = new Date();
            d.setTime(d.getTime() + (expire_min*60*1000));
            var expires = "expires="+ d.toUTCString();
            // var expires = "expires=" + d.toLocaleString();
            document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
        },
    },
    mounted() {
        this.sorryfornotwork();
    },
};
</script>

<style scoped>
#container {
    width: 100%;
    min-height: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3);
    display: flex;
    justify-content: center;
    align-items: center;
    /* background-image: url("~@/assets/login_background.jpg"); */
    background-color: rgba(0,0,0,0.5);
}
.surface {
    height: 500px;
    width: 900px;
    background-color: rgba(0,0,0,0.3);
    display: flex;
    position: relative;
    overflow: hidden;
    align-items: center;
    /* box-shadow: 0 5px 5px #ddd; */
}
.surface > .left {
    height: 400px;
    width: 500px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
}
.surface > .left > div {
    width: 100%;
    cursor: pointer;
}
.surface > .left > div > div {
    width: 40%;
    font-size: 36px;
    font-weight: bold;
    text-align: right;
    margin: 10px;
}
.surface > .right {
    position: absolute;
    transform: rotate(70deg);
    width: 700px;
    height: 700px;
    top: -150px;
    left: 400px;
    background-color: #aaa;
}
.login-zone {
    height: 400px;
    width: 400px;
    transform: rotate(-70deg);
    position: absolute;
    top: 235px;
    left: 175px;
    display: flex;
    justify-content: center;
    align-items: center;
}
.login-zone > div {
    width: 80%;
    height: 80%;
}
.login-zone > div > div {
    margin-top: 40px;
}
</style>
