// Function to show error toast
function showErrorToast(message) {
    var toastEl = document.getElementById('error-toast');
    var bsToast = new bootstrap.Toast(toastEl);
    document.getElementById("error-toast-body").innerText = message;
    bsToast.show();
}

// Function to show success toast
function showSuccessToast(message) {
    var toastEl = document.getElementById('success-toast');
    var bsToast = new bootstrap.Toast(toastEl);
    document.getElementById("success-toast-body").innerText = message;
    bsToast.show();
}
