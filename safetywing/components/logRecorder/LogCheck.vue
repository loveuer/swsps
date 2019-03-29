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
                @change="logRecorderCheckDateChg"
                placeholder="选择日期">
            </el-date-picker>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";
import dateformat from "dateformat";

export default {
    data() {
        return {
            checkDate: this.$route.params.date,
        };
    },
    methods: {
        logRecorderCheckDateChg: function() {
            if (!this.checkDate) {
                return;
            };
            let checkedDate = dateformat(this.checkDate, "yyyy-mm-dd");
            this.$router.push("/works/logRecorder/check/" + checkedDate);
            // axios
            // get some
            console.log("chged");
        },
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/logRecorder/check/" + this.$route.params.date;
        // axios
    },
    watch: {
        '$route.params.date': function(newval, oldval) {
            this.checkDate = this.$route.params.date;
            // axios
        }
    }
};
</script>

