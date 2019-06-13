<template>
    <div id="container">
        <el-table
            :data="sps"
            style="font-size:14px;width:100%;cursor:pointer"
            :row-class-name="tableRowClassName"
            @row-click="handleRowClick"
        >
            <el-table-column type="expand">
                <template slot-scope="props">
                    <el-form label-position="left" class="demo-table-expand" label-width="100px">
                        <el-form-item label="原模拟机" style="height:30px;">
                            <span>{{ props.row.orgsim }}</span>
                        </el-form-item>
                        <el-form-item label="位　　置" style="height:30px;">
                            <span>{{ props.row.pos }}</span>
                        </el-form-item>
                        <el-form-item label="状　　态" style="height:30px;">
                            <span>{{ props.row.status }}</span>
                        </el-form-item>
                        <el-form-item label="备　　注" style="height:30px;">
                            <span>{{ props.row.comment }}</span>
                        </el-form-item>
                    </el-form>
                </template>
            </el-table-column>
            <el-table-column prop="name" label="Name" min-width="300"></el-table-column>
            <el-table-column prop="pn" label="P/N" min-width="250"></el-table-column>
            <el-table-column prop="sn" label="S/N" min-width="200"></el-table-column>
            <el-table-column prop="nowsim" label="模拟机" min-width="75"></el-table-column>
            <el-table-column prop="amount" label="数量" min-width="75"></el-table-column>
        </el-table>
    </div>
</template>

<script>
export default {
    props: {
        sps: Array,
    },
    methods: {
        tableRowClassName: function(row) {
            if (row.row.status === '废弃') {
                return 'danger-row';
            };
            return '';
        },
        handleRowClick: function(row) {
            this.$router.push(`/onesp/${row.id}`);
        },
    }
}
</script>

<style>
.el-table .danger-row {
    background: oldlace;
}
.el-table .success-row {
    background: #5ca833;
}
</style>
