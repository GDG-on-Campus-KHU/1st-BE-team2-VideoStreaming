document.addEventListener('DOMContentLoaded', function() {
    loadVideoList();

    const uploadButton = document.getElementById('uploadButton');
    const videoUploadInput = document.getElementById('videoUploadInput');

    // 버튼 클릭 시 파일 선택 창 열기
    uploadButton.addEventListener('click', function(){
        videoUploadInput.click();
    });

    // 파일 선택 시 서버로 업로드
    videoUploadInput.addEventListener("change", function(){
        const file = videoUploadInput.files[0];
        if (file){
            uploadVideo(file);
        }
    });
});

function loadVideoList() {
    fetch('/videos')
        .then(response => response.json())
        .then(videos => {
            const videoList = document.getElementById('videoList');
            videos.forEach(video => {
                const li = document.createElement('li');
                li.textContent = video;
                li.addEventListener('click', function() {
                    playVideo(video);
                });
                videoList.appendChild(li);
            });
        })
        .catch(error => console.error('Error loading video list:', error));
}

function playVideo(videoName) {
    console.log('Playing video:', videoName);
    const videoPlayer = document.getElementById('videoPlayer');
    videoPlayer.src = `/stream?video=${encodeURIComponent(videoName)}`;
    videoPlayer.play().catch(e => console.error('Error playing video:', e));
}

function uploadVideo(file){
    const formData = new FormData();
    formData.append('video', file);
    
    fetch('/upload',{
        method: 'POST',
        body: formData,
    })
    .then(response => {
        if (response.ok){
            console.log("Video Uploaded successfully");
            loadVideoList(); // 업로드 후 동영상 리스트 refresh
        } else {
            console.error('Error uploading video');
        }
    })
    .catch(error => console.error("Error uploading video", error));
}