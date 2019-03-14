<template>
    <div>
        <el-form label-position="left" class="demo-table-expand" label-width="150px">
            <el-form-item label="ID:">
                <span>{{ this.$store.state.spsOne.id }}</span>
            </el-form-item>
            <el-form-item label="名字:">
                <span>{{ this.$store.state.spsOne.name }}</span>
            </el-form-item>
            <el-form-item label="P/N:">
                <span>{{ this.$store.state.spsOne.pn }}</span>
            </el-form-item>
            <el-form-item label="S/N:">
                <span>{{ this.$store.state.spsOne.sn }}</span>
            </el-form-item>
            <el-form-item label="模拟机:">
                <span>{{ this.$store.state.spsOne.sim }}</span>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            sp: this.$store.state.spsOne,
        };
    },
};
</script>

<style scoped>
.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
