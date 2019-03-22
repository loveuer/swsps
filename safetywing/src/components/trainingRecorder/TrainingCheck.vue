<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="margin-left:30px;margin-top:30px;">
            <span class="demonstration">日期: </span>
            <el-date-picker
                v-model="checkDate"
                type="date"
                @change="trainingRecorderCheckDateChg"
                placeholder="选择日期">
            </el-date-picker>
        </div>
        <!-- content -->
        <div>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import dateform from "dateformat";

export default {
    data() {
        return {
            checkDate: this.$route.params.date,
        }   
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/trainingRecorder/check/" + this.$route.params.date;
        this.checkDate = this.$route.params.date;
        // axios
    },
    methods: {
        trainingRecorderCheckDateChg: function() {
            if (!this.checkDate) {
                return;
            };
            let checkedDate = dateform(this.checkDate, "yyyy-mm-dd");
            this.$router.push("/works/trainingRecorder/check/" + checkedDate);
        },
    },
    watch: {
        '$route.params.date': function(newval, oldval) {
            this.checkDate = this.$route.params.date;
            // axios
        },
    },
    
};
</script>

<style scoped>

</style>
