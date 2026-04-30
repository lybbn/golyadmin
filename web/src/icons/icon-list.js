const svgModules = import.meta.glob('./svg/*.svg', { eager: true, query: '?url', import: 'default' })

function getIconList() {
    let iconList = []
    Object.keys(svgModules).forEach(e => {
        const cname = e.split('/').pop()?.split('.')[0]
        iconList.push("lyicon-" + cname)
    })
    return iconList
}

export {
    getIconList
}
