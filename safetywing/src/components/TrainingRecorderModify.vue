<template>
    <div id="t-modify">
        <el-row>
            <el-date-picker v-model="modifyDate" type="date" placeholder="Pick a day" @change="getScanedListByDate"> </el-date-picker>
        </el-row>
    </div>
</template>

<script>
export default {
    data() {
        return {
            modifyDate: "",
        };
    },
    methods: {
        getScanedListByDate: function() {
            
        }
    }
};
</script>

<style scoped>
#t-modify{
    margin-left: 20px;
    margin-top: 20px;
}
</style>
