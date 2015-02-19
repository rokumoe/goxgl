// +build windows

package context

//#cgo windows	LDFLAGS: -lopengl32 -lgdi32
import "C"

/*
enum KEYCODE {
	KEYCODE_LEFT,
	KEYCODE_RIGHT,
	KEYCODE_UP,
	KEYCODE_DOWN,
	KEYCODE_UNKNOWN
};

extern void onInit();
extern void onExit();
extern void onSize(int x, int y);
extern void onDraw();
extern void onKeyDown(int kc);
extern void onKeyUp(int kc);
*/
import "C"

/*
#include <Windows.h>

HWND _hWnd = NULL;
HDC _hDC = NULL;
HGLRC _hGLRC = NULL;
int _inited = 0;

LRESULT WINAPI _wndProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
	PAINTSTRUCT ps;

	switch (uMsg) {
	case WM_PAINT:
		BeginPaint(hWnd, &ps);
		onDraw();
		EndPaint(hWnd, &ps);
		break;
	case WM_SIZE:
		onSize(LOWORD(lParam), HIWORD(lParam));
		break;
	case WM_KEYDOWN:
		switch (wParam) {
		case VK_LEFT:
			onKeyDown(KEYCODE_LEFT);
			break;
		case VK_RIGHT:
			onKeyDown(KEYCODE_RIGHT);
			break;
		case VK_UP:
			onKeyDown(KEYCODE_UP);
			break;
		case VK_DOWN:
			onKeyDown(KEYCODE_DOWN);
			break;
		default:
			onKeyDown(KEYCODE_UNKNOWN);
		}
		break;
	case WM_KEYUP:
		switch (wParam) {
		case VK_LEFT:
			onKeyUp(KEYCODE_LEFT);
			break;
		case VK_RIGHT:
			onKeyUp(KEYCODE_RIGHT);
			break;
		case VK_UP:
			onKeyUp(KEYCODE_UP);
			break;
		case VK_DOWN:
			onKeyUp(KEYCODE_DOWN);
			break;
		default:
			onKeyUp(KEYCODE_UNKNOWN);
		}
		break;
	case WM_DESTROY:
		if (hWnd == _hWnd) {
			if (_inited) {
				onExit();
				_inited = 0;
			}
			wglDeleteContext(_hGLRC);
			_hGLRC = NULL;
			wglMakeCurrent(NULL, NULL);
			ReleaseDC(_hWnd, _hDC);
			_hDC = NULL;
			_hWnd = NULL;
		}
		PostQuitMessage(0);
		break;
	default:
		return DefWindowProc(hWnd, uMsg, wParam, lParam);
	}
	return 0;
}

int initDisplay(const char * name, int x, int y, int w, int h)
{
	HINSTANCE instance;
	WNDCLASS wc;
	PIXELFORMATDESCRIPTOR pfd;
	int pf;

	instance = (HINSTANCE)GetModuleHandle(NULL);
	wc.style = CS_OWNDC;
	wc.lpfnWndProc = _wndProc;
	wc.cbClsExtra = 0;
	wc.cbWndExtra = 0;
	wc.hInstance = instance;
	wc.hIcon = LoadIcon(NULL, IDI_APPLICATION);
	wc.hCursor = LoadCursor(NULL, IDC_ARROW);
	wc.hbrBackground = NULL;
	wc.lpszMenuName = NULL;
	wc.lpszClassName = name;
	if (!RegisterClass(&wc)) {
		return 1;
	}
	if (_hWnd == NULL) {
		_hWnd = CreateWindow(name, name,
							 WS_OVERLAPPEDWINDOW | WS_CLIPCHILDREN | WS_CLIPSIBLINGS,
							 x, y, w, h, NULL, NULL, instance, NULL);
		if (_hWnd == NULL) {
			return 2;
		}
	}
	if (_hDC == NULL) {
		_hDC = GetDC(_hWnd);
		if (_hDC == NULL) {
			return 3;
		}
	}
	if (_hGLRC == NULL) {
		ZeroMemory(&pfd, sizeof(pfd));
		pfd.nSize = sizeof(pfd);
		pfd.nVersion = 1;
		pfd.dwFlags = PFD_DRAW_TO_WINDOW | PFD_SUPPORT_OPENGL | PFD_DOUBLEBUFFER;
		pfd.iPixelType = PFD_TYPE_RGBA;
		pfd.cColorBits = 32;
		pfd.iLayerType = PFD_MAIN_PLANE;
		pf = ChoosePixelFormat(_hDC, &pfd);
		if (pf == 0) {
			return 4;
		}
		if (!SetPixelFormat(_hDC, pf, &pfd)) {
			return 5;
		}
		_hGLRC = wglCreateContext(_hDC);
	}
	if (!wglMakeCurrent(_hDC, _hGLRC)) {
		return 6;
	}
	onInit();
	_inited = 1;
	ShowWindow(_hWnd, SW_NORMAL);
	return 0;
}

void mainLoop()
{
	MSG msg;

	while (GetMessage(&msg, _hWnd, 0, 0)) {
		TranslateMessage(&msg);
		DispatchMessage(&msg);
	}
}

void requestRedraw()
{
	InvalidateRect(_hWnd, NULL, FALSE);
}

void swapBuffer()
{
	SwapBuffers(_hDC);
}
*/
import "C"

func InitDisplay(name string, x, y, w, h int) int {
	s := C.CString(name)
	return int(C.initDisplay(s, C.int(x), C.int(y), C.int(w), C.int(h)))
}

func MainLoop() {
	C.mainLoop()
}

func RequsetRedraw() {
	C.requestRedraw()
}

func SwapBuffer() {
	C.swapBuffer()
}
