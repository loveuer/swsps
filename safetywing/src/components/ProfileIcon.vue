<template>
    <div id="profile-icon" v-bind:title="this.$store.state.user.realname">
        <div v-show="this.$store.state.user.profileIcon !== ''" id="img">
            <img v-bind:src="this.$store.state.user.profileIcon">
        </div>
        <div v-show="this.$store.state.user.profileIcon === ''" id="font">{{ this.$store.state.user.realname.slice(0,1) }}</div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            color: ["#0fa2a9", "#949217", "#c1a7b0", "#7d3865", "#d7743b"],
        }
    },
    methods: {
        getProfile: function() {
            this.$http.get("/api/user/" + this.$store.state.user.id)
                .then(resp => {
                    if(resp.data["Msg"] === "200"){
                        let user = JSON.parse(resp.data["Val"]);
                        let newUser = {
                            id: user.id,
                            username: user.username,
                            realname: user.realname,
                            profileIcon: user.img,
                        }
                        this.$store.commit("chgUser", newUser);
                        
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
