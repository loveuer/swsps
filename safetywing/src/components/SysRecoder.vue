<template>
    <div>
        <el-row>this is test page now</el-row>
        <el-row>
            <div v-for="(v, index) of vs" :key="index">
                <p>{{ v }}</p>
            </div>

        </el-row>
    </div>
</template>

<script>
export default {
    computed: {
        vs: function() {
            return [
                1,2,3,3,9,1,1,9,
            ]
        },
    },
}
</script>

<style scoped>

</style>

