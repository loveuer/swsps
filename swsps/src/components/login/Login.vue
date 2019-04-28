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
                                <el-button type="primary" @click="dologin">登录</el-button>
                                <el-button>注册</el-button>
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
        };
    },
    methods: {
        meetHome() {
            this.$router.push('/');
        },
        dologin() {
            if (!this.username || !this.password) {
                this.$message({
                    showClose: true,
                    message: '用户名和密码不能为空',
                    type: 'warning',
                });
                this.password = '';
                return;
            }
        },
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
