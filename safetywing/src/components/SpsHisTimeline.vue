<template>
    <el-timeline :reverse="false">
        <el-timeline-item
            v-for="(hs, index) of histories"
            :key="index"
            :timestamp="hs.date"
        >
            <el-table
                :data="hs.his"
                style="width: 100%">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" class="demo-table-expand">
                            <el-form-item label="模拟机:">
                                <span>{{ props.row.oldsim + " ->" + props.row.newsim }}</span>
                            </el-form-item>
                            <el-form-item label="位置:">
                                <span>{{ props.row.oldpos + " ->" + props.row.newpos }}</span>
                            </el-form-item>
                            <el-form-item label="状态:">
                                <span>{{ props.row.oldstatus + " ->" + props.row.newstatus }}</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                <el-table-column
                    label="时间"
                    prop="time">
                </el-table-column>
                <el-table-column
                    label="人员"
                    prop="auth">
                </el-table-column>
                <el-table-column
                    label="简述"
                    prop="short">
                </el-table-column>
            </el-table>
        </el-timeline-item>
        
    </el-timeline>
   
</template>

<style>
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

<script>
export default {
    data() {
      return {
        histories: [],
        otherdata: {},
        };
    },
}
</script>