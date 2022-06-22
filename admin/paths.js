const fs = require('fs')

fs.readFile('./build/index.html', 'utf8', (err, data) => {
  data = data.replaceAll('"/static', '"static')
  //fs.unlink("./build/index.html", () => { })
  fs.writeFileSync("./build/index.html", data)
  
})
