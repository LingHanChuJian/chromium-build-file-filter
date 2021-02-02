# chromium-build-file-filter
chromium 打包文件过滤

### chromium build
chromium 打包 有用文件目录如下

```TEXT
chromium
├─ 89.0.4335.0.manifest
├─ chrome.dll
├─ chrome.exe
├─ chrome_100_percent.pak
├─ chrome_200_percent.pak
├─ chrome_elf.dll
├─ headless_lib.pak
├─ icudtl.dat
├─ libEGL.dll
├─ libGLESv2.dll
├─ locales
│  ├─ am.pak
│  ├─ am.pak.info
│  ├─ ar.pak
│  ├─ ar.pak.info
│  ├─ bg.pak
│  ├─ bg.pak.info
│  ├─ bn.pak
│  ├─ bn.pak.info
│  ├─ ca.pak
│  ├─ ca.pak.info
│  ├─ cs.pak
│  ├─ cs.pak.info
│  ├─ da.pak
│  ├─ da.pak.info
│  ├─ de.pak
│  ├─ de.pak.info
│  ├─ el.pak
│  ├─ el.pak.info
│  ├─ en-GB.pak
│  ├─ en-GB.pak.info
│  ├─ en-US.pak
│  ├─ en-US.pak.info
│  ├─ es-419.pak
│  ├─ es-419.pak.info
│  ├─ es.pak
│  ├─ es.pak.info
│  ├─ et.pak
│  ├─ et.pak.info
│  ├─ fa.pak
│  ├─ fa.pak.info
│  ├─ fi.pak
│  ├─ fi.pak.info
│  ├─ fil.pak
│  ├─ fil.pak.info
│  ├─ fr.pak
│  ├─ fr.pak.info
│  ├─ gu.pak
│  ├─ gu.pak.info
│  ├─ he.pak
│  ├─ he.pak.info
│  ├─ hi.pak
│  ├─ hi.pak.info
│  ├─ hr.pak
│  ├─ hr.pak.info
│  ├─ hu.pak
│  ├─ hu.pak.info
│  ├─ id.pak
│  ├─ id.pak.info
│  ├─ it.pak
│  ├─ it.pak.info
│  ├─ ja.pak
│  ├─ ja.pak.info
│  ├─ kn.pak
│  ├─ kn.pak.info
│  ├─ ko.pak
│  ├─ ko.pak.info
│  ├─ lt.pak
│  ├─ lt.pak.info
│  ├─ lv.pak
│  ├─ lv.pak.info
│  ├─ ml.pak
│  ├─ ml.pak.info
│  ├─ mr.pak
│  ├─ mr.pak.info
│  ├─ ms.pak
│  ├─ ms.pak.info
│  ├─ nb.pak
│  ├─ nb.pak.info
│  ├─ nl.pak
│  ├─ nl.pak.info
│  ├─ pl.pak
│  ├─ pl.pak.info
│  ├─ pt-BR.pak
│  ├─ pt-BR.pak.info
│  ├─ pt-PT.pak
│  ├─ pt-PT.pak.info
│  ├─ ro.pak
│  ├─ ro.pak.info
│  ├─ ru.pak
│  ├─ ru.pak.info
│  ├─ sk.pak
│  ├─ sk.pak.info
│  ├─ sl.pak
│  ├─ sl.pak.info
│  ├─ sr.pak
│  ├─ sr.pak.info
│  ├─ sv.pak
│  ├─ sv.pak.info
│  ├─ sw.pak
│  ├─ sw.pak.info
│  ├─ ta.pak
│  ├─ ta.pak.info
│  ├─ te.pak
│  ├─ te.pak.info
│  ├─ th.pak
│  ├─ th.pak.info
│  ├─ tr.pak
│  ├─ tr.pak.info
│  ├─ uk.pak
│  ├─ uk.pak.info
│  ├─ vi.pak
│  ├─ vi.pak.info
│  ├─ zh-CN.pak
│  ├─ zh-CN.pak.info
│  ├─ zh-TW.pak
│  └─ zh-TW.pak.info
├─ resources.pak
├─ snapshot_blob.bin
├─ swiftshader
│  ├─ libEGL.dll
│  └─ libGLESv2.dll
└─ v8_context_snapshot.bin
```
