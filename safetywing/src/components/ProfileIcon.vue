<template>
    <div id="profile-icon" v-bind:title="realname">
        <div v-show="img" id="img">
            <img src="">
        </div>
        <div v-show="!img" id="font"></div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            img: false,
            xing: "",
            realname: "",
            color: ["#0fa2a9", "#949217", "#c1a7b0", "#7d3865", "#d7743b"],
        }
    },
    methods: {
        getProfile: function() {
            this.$http.get("/api/user/" + this.$user.id)
                .then(resp => {
                    if(resp.data["Msg"] === "200"){
                        let user = JSON.parse(resp.data["Val"]);
                        this.$user.username = user.username;
                        this.$user.realname = user.realname;
                        this.$user.profileIcon = user.img;
                        this.xing = user.realname.slice(0,1);
                        this.realname = user.realname;
                        
                        document.querySelector("#img img").src = user.img;

                        if(user.img !== "") {
                            this.img = !this.img;
                        }
                    };
                })
                .catch(error => { console.log(error)} );
        },
    },
    mounted() {

    },
}
</script>

<style scoped>
#profile-icon{
    min-width: 100%;
    min-height: 100%;
    border-radius: 50%;

}
#font{

}
#img{
    min-width: 100%;
    min-height: 100%;
    display: inline-block;
    overflow: hidden;
}
#img > img {
    border-radius: 50%;
    height: auto;
    max-width: 100%;
}
</style>
