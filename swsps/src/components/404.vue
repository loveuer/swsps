<template>
    <div id="container">
        <h1>oops!~~~</h1>
        <el-button @click="meetHome">返回主页</el-button>
    </div>
</template>

<script>
export default {
    methods: {
        meetHome() {
            this.$router.push("/");
        },
    },
}
</script>
