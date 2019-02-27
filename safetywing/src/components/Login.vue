<template>
    <div class="container">
        <div class="login-zone">
            <div class="login">
                <loginform />
            </div>
        </div>
    </div>
</template>

<script>
import loginform from "./LoginForm.vue";

export default {
    data() {
        return {
            
        }
    },
    components: {
        loginform,
    },
    mounted() {

    },
    methods: {
        
    }
};
</script>

<style scoped>
.container{
    position: absolute;
    height: 100%;
    width: 100%;
    background-image: url("~@/assets/img/runway.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
}
.login-zone{
    height: 100%;
    width: 100%;
    background-color: rgba(0,0,0,0.3);
    filter:Alpha(opacity=50);
    display: flex;
    justify-content: center;
    align-items: center;
}
.login{
    height: 270px;
    width: 340px;
    border: 2px solid #fff;
    padding: 20px;
    border-radius: 20px;
    background-color: rgba(222,222,222,0.7);
}
.login:hover{
}
</style>
