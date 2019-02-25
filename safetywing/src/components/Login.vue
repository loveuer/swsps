<template>
    <div class="container">
        <div class="login-zone">
            <a-form
                :form="form"
                @submit="handleSubmit"
            >
                <a-form-item
                    label="username"
                    :label-col="{ span: 5 }"
                    :wrapper-col="{ span: 12 }"
                >
                    <a-input
                        v-decorator="[
                            'note',
                            {rules: [{ required: true, message: '请输入用户名:'}]}
                        ]"
                    />
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            formLayout: "horizontal",
            form: this.$form.createForm(this),
        };
    },
    mounted() {
        document.querySelector(".container").style.height = window.innerHeight + "px";
    },
    methods: {
        handleSubmit: (e) => {

        },
    }
};
</script>

<style scoped>
.container{
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    background-image: url("~@/assets/img/runway.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
    opacity: 0.5;
}
.login-zone{
    height: 250px;
    width: 450px;
    border: 2px solid #fff;
    border-radius: 15px;
    background-color: rgba(240, 240, 240, 0.5);
}
.login-zone:hover{
    border: 2px solid #0082E4;
}
</style>
