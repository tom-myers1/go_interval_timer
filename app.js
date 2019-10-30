const go = new Go()
WebAssembly.instantiateStreaming(fetch('app.wasm', {
  headers: {
    'Content-Type': 'application/wasm'
  }
}), go.importObject).then(result => go.run(result.instance))
