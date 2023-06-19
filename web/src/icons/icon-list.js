const modules = require.context('./svg', false, /\.svg$/);

function getIconList() {
    let iconList = []
    modules.keys().forEach(e => {
        const cname = e.split('/').pop()?.split('.')[0]//根据路径截取name文件名（去除后缀和前面目录）
        iconList.push("lyicon-"+cname)
    })
    return iconList
}

export {
    getIconList
}
