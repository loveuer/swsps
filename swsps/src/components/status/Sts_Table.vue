<template>
    <div>
        <el-table
            ref="filterTable"
            :data="showingList"
            style="width:100%"
            v-loading="false"
        >
            <el-table-column
                prop="name"
                label="Name"
                width="200"
            >
            </el-table-column>
            <el-table-column
                prop="pn"
                label="P/N"
                width="200"
            >
            </el-table-column>
            <el-table-column
                prop="sn"
                label="S/N"
                width="200"
            >
            </el-table-column>
            <el-table-column
                prop="orgsim"
                label="原模拟机"
                width="120"
            >
            </el-table-column>
            <el-table-column
                prop="status"
                label="状态"
                width="100"
                :filters="[{ text: '使用', value: '使用' }, { text: '维修', value: '维修' }, { text: '废弃', value: '废弃' }]"
                :filter-method="fliterStatus"
            >
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
export default {
    props: {
        showingMap: Object,
    },
    methods: {
        fliterStatus(value, row) {
            return row.status === value;
        },
    },
    computed: {
        showingList: function() {
            return [].concat(this.showingMap.using).concat(this.showingMap.fixing).concat(this.showingMap.abandon);
        },
    },
};
</script>

<style scoped>

</style>