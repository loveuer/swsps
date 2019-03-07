<template>
    <div id="profile-icon" v-bind:title="this.$store.state.user.realname">
        <div v-show="this.$store.state.user.profile_icon !== ''" id="img">
            <img v-bind:src="this.$store.state.user.profile_icon">
        </div>
        <div v-show="this.$store.state.user.profile_icon === ''" id="font">{{ this.$store.state.user.realname.slice(0,1) }}</div>
    </div>
</template>

<script>
export default {
    data() {
        return {
        }
    },
    methods: {
        getProfile: function() {
            this.$http.get("/api/user/" + this.$store.state.user.username)
                .then(resp => {
                    if(resp.status === 200){
                        let user = JSON.parse(resp.data["val"]);
                        let newUser = {
                            id: user.id,
                            username: user.username,
                            realname: user.realname,
                            profile_icon: "http://127.0.0.1:8000/static/img/" + user.profile_icon,
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
