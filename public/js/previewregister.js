const previewImage = (event) => {
    const array_of_allowed_files = ['png', 'jpeg', 'jpg'];
    var tweet_image = document.getElementById("image");
    const file_extension=tweet_image.value.slice(((tweet_image.value.lastIndexOf('.') - 1)>>>0) + 2);
    // console.log(file_extension)
    var imgsize = Math.round(tweet_image.files[0].size / 1048576);
    console.log(imgsize)
    if (imgsize > 2) {
        document.getElementById("err2").innerHTML = "upload image whose size is less than 2mb";
        document.getElementById("demo").disabled = true;
    }
    else {
        if (array_of_allowed_files.includes(file_extension)) {
            document.getElementById("err2").innerHTML = "";
            /**
            * Get the selected files.
            */
            const imageFiles = event.target.files;
            /**
            * Count the number of files selected.
            */
            const imageFilesLength = imageFiles.length;
            /**
            * If at least one image is selected, then proceed to display the preview.
            */
            if (imageFilesLength > 0) {
                /**
                 * Get the image path.
                 */
                const imageSrc = URL.createObjectURL(imageFiles[0]);
                /**
                 * Select the image preview element.
                 */
                const imagePreviewElement = document.querySelector("#preview");
                /**
                 * Assign the path to the image preview element.
                 */
                imagePreviewElement.src = imageSrc;
                /**
                 * Show the element by changing the display value to "block".
                 */
                imagePreviewElement.style.display = "block";
            }
            document.getElementById("demo").disabled = false;
        } else {
            document.getElementById("err2").innerHTML = "upload only jpg,jpeg and png formate files";
            document.getElementById("demo").disabled = true;
        }

    }

};