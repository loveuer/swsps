<template>
    <div class="container" :style="{display:ifshowStr}" @keydown.esc="quit" tabindex="0" @click="quit">
        <img :src="imgSrc" @click="openImg">
    </div>
</template>

<script>
export default {
    data() {
        return {
            imgSrc: "",
            ifshow: false,
        };
    },
    methods: {
        openImg: function() {
            window.open(this.imgSrc, "_blank");
        },
        quit: function() {
            this.ifshow = false;
        },
    },
    computed: {
        ifshowStr: function() {
            if (this.ifshow) {
                return 'block';
            } else {
                return 'none';
            };
        },
    }
};
</script>


<style scoped>
.container{
    width:100%;
    height:100%;
    position: fixed;
    z-index: 999;
    top: 0;
    left: 0;
    background-color: rgba(0,0,0,0.8);
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    /* animation-name: lightboxAnimation;
    animation-fill-mode: forwards;
    animation-duration: .1s; */
}
@keyframes lightboxAnimation{
    from { width:0%; height:0%; }
    to { width:100%; height:100%; }
}
.container > img {
    max-width: 90%;
    max-height: 80%;
    margin-top: 2%;
    cursor: zoom-in;
}
</style>
