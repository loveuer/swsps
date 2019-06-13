<template>
    <div id="container">
        <el-table
            :data="folded_sps"
            style="font-size:14px;width:100%;"
        >
            <el-table-column type="expand">
                <template slot-scope="props">
                    <!-- <el-table :data="props.row.detail" style="font-size:0.8rem;cursor:pointer;overflow-x:hidden;" :show-header='false'>
                        <el-table-column prop="name" label="Name" width="400" :show-overflow-tooltip="true"></el-table-column>
                        <el-table-column prop="pn" label="P/N" width="300" :show-overflow-tooltip="true"></el-table-column>
                        <el-table-column prop="sn" label="S/N" width="300" :show-overflow-tooltip="true"></el-table-column>
                        <el-table-column prop="nowsim" label="模拟机" width="100"></el-table-column>
                        <el-table-column prop="amount" label="数量" width="100"></el-table-column>
                    </el-table> -->
                    <div v-for="(item, index) of props.row.detail" :key="index" class="foldedtable-inner">
                        <div style="width:400px;margin-left:8px;">{{ item.name }}</div>
                        <div style="width:300px;margin-left:38px;">{{ item.pn }}</div>
                        <div style="width:300px;margin-left:38px;">{{ item.sn }}</div>
                        <div style="width:100px;margin-left:38px;">{{ item.nowsim }}</div>
                        <div style="width:100px;margin-left:38px;">{{ item.amount }}</div>
                    </div>
                </template>
            </el-table-column>
            <el-table-column prop="name" label="Name" width="400" :show-overflow-tooltip="true"></el-table-column>
            <el-table-column prop="pn" label="P/N" width="300" :show-overflow-tooltip="true"></el-table-column>
            <el-table-column prop="sn" label="S/N" width="300" :show-overflow-tooltip="true"></el-table-column>
            <el-table-column prop="nowsim" label="模拟机" width="100"></el-table-column>
            <el-table-column prop="amount" label="数量" width="100"></el-table-column>
        </el-table>
        <br><br>
        <el-button @click="show_folded_sp">先看看</el-button>
    </div>
</template>

<script>
export default {
    props: {
        sps: Array,
    },
    computed: {
        folded_sps: function() {
            let fp = {};
            let rfp = [];

            for (let sp of this.sps) {
                if (sp.pn === 'None') {
                    rfp.push({name:sp.name, pn:'None', amount:sp.amount, detail:[sp]});
                    continue;
                };
                if (typeof(fp[sp.pn]) === "undefined") {
                    let spA = 0;
                    try {
                        spA = parseInt(sp.amount);
                    } catch (error) {
                        console.log(error);
                    };
                    fp[sp.pn] = {name: sp.name, pn: sp.pn, amount: spA, detail: [sp,]};
                } else {
                    let addA = 0;
                    try {
                        addA = parseInt(sp.amount);
                    } catch (error) {
                        console.log(error);
                    };
                    fp[sp.pn]["amount"] += parseInt(sp.amount)
                    fp[sp.pn]["detail"].push(sp);
                };
            };
            for (let one in fp) {
                rfp.push(fp[one]);
            };
            return rfp;
        },
    },
    methods: {
        show_folded_sp: function(params) {
            console.log(this.folded_sps);  
        },
    },
}
</script>

<style scoped>
.table-title {
    width: 100%;
    display: flex;
    align-items: center;
}
.table-title > div {
    font-size: 0.9rem;
    margin-bottom: 5px;
    
}
.foldedtable-inner {
    display: flex;
}
.foldedtable-inner > div {

}
</style>
