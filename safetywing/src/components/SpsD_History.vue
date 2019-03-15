<template>
    <div style="margin-top:20px;">
        <sps-his-timeline ref="spshistimeline"></sps-his-timeline>
    </div>
</template>

<script>
import spsHisTimeline from "./SpsHisTimeline.vue";

export default {
    mounted() {
        console.log("sp id: " + this.$store.state.spsOne.id);
        this.$refs.spshistimeline.histories = this.$store.state.mockhis;
    },
    components: {
        'sps-his-timeline': spsHisTimeline,
    }
};
</script>
