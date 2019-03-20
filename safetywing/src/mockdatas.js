const mockdata = {
    mockhis: [
        {spid:105,date:'2019-02-20T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-02-20T00:00:00Z',time:'0000-01-01T12:36:49Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-02-20T00:00:00Z',time:'0000-01-01T12:35:15Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-02-19T00:00:00Z',time:'0000-01-01T00:30:35Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-01-29T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-01-29T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-01-29T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-01-29T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-02-20T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
        {spid:105,date:'2019-02-20T00:00:00Z',time:'0000-01-01T15:32:36Z',last_sim:'5015',next_sim:'5015',last_pos:'a-3-1',next_pos:'a-3-1',last_sts:'备件',next_sts:'备件',amount_chg:-1,auth:'李跃',comment:'0000-01-01T15:32:36Z',short:null},
    ],
    mockSearchSps: {
        'byname': [
            {'id':2,'name':'du','pn':'221','sn':'0371','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
            {'id':22,'name':'du','pn':'221','sn':'0372','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
            {'id':222,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
        ],
        'bypn': [
            {'id':33,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
            {'id':333,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
        ],
        'bysn': [
            {'id':4,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
            {'id':44,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
            {'id':444,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
            {'id':8,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
        ],
        'bycomment': [
            {'id':11,'name':'du','pn':'221','sn':'0373','orgsim':'5978','nowsim':'5989','pos':'a-1-3','amount':1, 'status':'备件','limit':-1,'is_consumable':false,'comment':'dudu','img1':'none','img2':'none'},
        ]
    }
};

export default mockdata;
