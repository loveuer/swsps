<template>
    <div>
        <el-row>
            this is modify sps
        </el-row>
        <div slot="footer" class="dialog-footer" :style="{marginTop:'30px'}">
            <el-button type="primary" @click="modifySPS">确认</el-button>
        </div>
    </div>
</template>

<script>
export default {
    methods: {
        modifySPS: function() {
            console.log("modify sps comfirm clicked");
            // post this.$store.state.spsOne balabala
        },
    },
};
</script>
