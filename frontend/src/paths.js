const fs = require('fs')


fs.readFile('./build/index.html', 'utf8', (err, data) => {
  //data = data.replaceAll('"/static', '"static')
  /*let script = ""
  let style = ""
  let i = data.indexOf('script defer="defer" src="') + 'script defer="defer" src="'.length;
  while (data[i] != '"') {
    script += data[i]
    i++
  }
  i = data.indexOf('link href="', i) + 'link href="'.length;
  while (data[i] != '"') {
    style += data[i]
    i++
  }
  fs.readFile("./build/"+script, "utf-8", (err, scriptT) => {
    fs.readFile("./build/"+style, "utf8", (err, styleT) => {
      //data = data.replace('<script defer="defer" src="' + script + '"></script>', '<script type="text/javascript">' + scriptT.replace(/[\\"']/g, '\\$&').replace(/\u0000/g, '\\0') + '</script>')
      data = data.replace('<link href="' + style + '" rel="stylesheet">', '<style>' + styleT + '</style>')
    
    
      fs.readFile("./src/scripts.js", "utf8", (err, scripts) => {
        data = data.replaceAll("</body>", "</body>\n<script>\n"+scripts+"\n</script>");
        fs.writeFileSync("./build/index.html", data)
      });
    })
  })*/
  
})

