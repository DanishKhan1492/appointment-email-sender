let dropArea = document.getElementById("drop-area");
let fileInput = document.getElementById("file");
let loaderOverlay = document.getElementById("loader-overlay");

dropArea.addEventListener("click", () => {
    fileInput.click();
});

dropArea.addEventListener("dragenter", preventDefaults, false);
dropArea.addEventListener("dragover", preventDefaults, false);
dropArea.addEventListener("dragleave", preventDefaults, false);
dropArea.addEventListener("drop", handleDrop, false);

function preventDefaults(e) {
    e.preventDefault();
    e.stopPropagation();
}

function handleDrop(e) {
    let dt = e.dataTransfer;
    let files = dt.files;
    handleFiles(files);
}

function handleFiles(files) {
    if (files.length > 0) {
        console.log("Uploading", files[0].name);
        showLoader();
        // Create FormData object
        let formData = new FormData();
        formData.append("file", files[0]);

        // Send POST request
        fetch("/", {
            method: "POST",
            body: formData
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.text();
            })
            .then(htmlContent => {
                // Update document body with received HTML content
                document.getElementById("customer-container").innerHTML = htmlContent;
                hideLoader();
            })
            .catch(error => {
                console.error("Error:", error);
                hideLoader();
            });
    }
}

function showLoader() {
    loaderOverlay.classList.add("disable-pointer-events");
    loaderOverlay.style.display = "block";
}

function hideLoader() {
    loaderOverlay.classList.remove("disable-pointer-events");
    loaderOverlay.style.display = "none";
}
