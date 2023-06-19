import LyPictureMultipleUpload from '@/components/upload/mutiple-pictures'
import LyPictureSingleUpload from '@/components/upload/single-picture'
import LyFileMultipleUpload from '@/components/upload/lyfileUpload'
import LyTeditor from '@/components/teditorjs/lyTEditor'

export default {
    install(app) {
        app.component("ly-public-pictrue-multiple-upload", LyPictureMultipleUpload);
        app.component("ly-public-pictrue-single-upload", LyPictureSingleUpload);
        app.component("ly-public-file-multiple-upload", LyFileMultipleUpload);
        app.component("ly-public-teditor", LyTeditor);
    },
};