document.addEventListener('DOMContentLoaded', function() {
    loadVideoList();
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