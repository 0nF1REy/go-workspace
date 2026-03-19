## ⚠️ Problema com GPUs antigas (OpenGL)

Em alguns sistemas Linux, especialmente com GPUs Intel antigas, o jogo pode falhar com o erro:

```
glfw: GLX: Failed to create context: GLXBadFBConfig
```

### 🧠 Causa

A biblioteca utilizada (Ebiten) requer **OpenGL 3.3 ou superior**, porém algumas GPUs antigas suportam apenas **OpenGL 2.1**.

---

## 🛠️ Como resolver

### ✅ 1. Forçar renderização por software

```bash
LIBGL_ALWAYS_SOFTWARE=1 go run main.go
```

✔ Maior compatibilidade  
❌ Menor desempenho (usa CPU)

---

### ✅ 2. Forçar compatibilidade com OpenGL 3.3

```bash
LIBGL_ALWAYS_SOFTWARE=1 MESA_GL_VERSION_OVERRIDE=3.3 go run main.go
```

✔ Pode evitar o erro de inicialização  
❌ Ainda depende de renderização por software

---

## 🔍 Verificar versão do OpenGL

```bash
glxinfo | grep "OpenGL version"
```

Se o resultado for algo como:

```
OpenGL version string: 2.1
```

👉 Sua GPU não atende ao requisito mínimo
