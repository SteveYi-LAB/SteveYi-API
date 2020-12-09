<?php
$URL='https://docs.google.com/uc';
$ACTION='export=download';
$LINK='googleusercontent.com';
$google_file_id_num='24,48';
$CODE='Yi';

$id=$_GET['id'];

if(empty($id)){
    echo "Please type file id";
} else{
    //GET USERCODE
    $download_link='https://docs.google.com/uc?export=download&id='.urlencode($id);
    $ch = curl_init($download_link);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_HEADER, 1);
    $result = curl_exec($ch);
    preg_match_all('/^Set-Cookie:\s*([^;]*)/mi', $result, $matches);
    $cookies = array();
    foreach($matches[1] as $item) {
        parse_str($item, $cookie);
        $cookies = array_merge($cookies, $cookie);
    }
    $USERCODE=$matches[1][0];
    $USERCODE=str_replace("download_warning_","",$USERCODE);
    $USERCODE=str_replace($id,"",$USERCODE);
    $USERCODE=substr($USERCODE,0,-6);
    curl_close($ch);
    
    //POST&GET COOKIE(Second)
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $download_link.'&confirm=Yi');
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'GET');
    $headers = array();
    $headers[] = 'Cookie: download_warning_'.urlencode($USERCODE).'_'.urlencode($id).'=Yi; Domain=.docs.google.com; Path=/uc; Secure; HttpOnly';
    curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
    $result = curl_exec($ch);
    if (curl_errno($ch)) {
        echo 'Error:' . curl_error($ch);
    }
    $DIRECT_LINK=strchr($result, "https://");
    $exp = "/.*(?=\">)/";
    preg_match($exp, $DIRECT_LINK, $MATCHED_STR);
    curl_close($ch);
    echo $DIRECT_LINK[0];
}
