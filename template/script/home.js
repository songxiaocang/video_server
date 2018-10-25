$(document).read(function () {
    DEFAULT_COOKIE_EXPIRE_TIME = 30;

    uname = '';
    session='';
    uid='';
    currentVideo= null;
    listVideos = null;

    session = getCookie('session')
    uname = getCookie('username')

    initPage(function () {
        if (listVideos !== null){
            currentVideo = listVideos[0];
            selectVideo(listVideos[0]['id']);
        }

        $(".video-item").on('click',function () {

        });
        $(".del-video-button").on('click',function () {

        });
        $("#submit-comment").on('click',function () {

        });
    })
    //home page event registry
    $("#regbtn").on('click',function (e) {
        
    });

    $("#signinbtn").on('click',function (e) {

    });

    $("#signinhref").on('click',function () {
        $("#regsubmit").hide();
        $("#signsubmit").show();
    });

    $("#registerhref").on('click',function () {
        $("#regsubmit").show();
        $("#signsubmit").hide();
    })

    //userhome page event registry
    $("#uploadform").on('submit',function (e) {
        
    });

    $(".close").on('click',function (e) {
        
    });

    $(".logout").on('click',function (e) {
        //清空cookie
    });

});
/* 同步操作 */
function initPage(callback) {



}

function setCookie(cname, cvalue, exmin){}

function getCookie(cname){}

//DOM operations
function selectVideo(vid) {
    
}

function refreshComments(vid){}

//弹出框
function popupNotificationMsg(msg) {
    
}

function popupErrorMsg(msg) {
    
}

//渲染评论
function htmlCommentListElement(cid, author, content){}

function htmlVideoListElement(vid,name,ctime) {
    
}

/* 异步操作 */
//Async ajax methods
//user operations
function registerUser(callback) {
    
}

function signinUser(callback){

}


function getUserId(callback){

}

//video operations
function createVideo(vname,callback){

}

function listAllVideos(callback){}

function deleteVideo(vid,callback){}


//comments operations
function postComment(vid,content,callback) {
    
}

function listAllComments(vid,callback) {
    
}