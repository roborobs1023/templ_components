package generator

import "fmt"

type GradientColor struct {
	from  string
	stops string
	to    string
}

var gradientColors = map[string]GradientColor{
	"from-transparent": {
		from:  "transparent",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(0, 0, 0, 0))",
	},
	"from-current": {
		from:  "currentColor",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(255, 255, 255, 0))",
	},
	"from-black": {
		from:  "var(--color-black)",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(0, 0, 0, 0))",
	},
	"from-white": {
		from:  "var(--color-white)",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(255, 255, 255, 0))",
	},
	"from-gray-50": {
		from:  "#f9fafb",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(249, 250, 251, 0))",
	},
	"from-gray-100": {
		from:  "#f3f4f6",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(243, 244, 246, 0))",
	},
	"from-gray-200": {
		from:  "#e5e7eb",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(229, 231, 235, 0))",
	},
	"from-gray-300": {
		from:  "#d1d5db",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(209, 213, 219, 0))",
	},
	"from-gray-400": {
		from:  "#9ca3af",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(156, 163, 175, 0))",
	},
	"from-gray-500": {
		from:  "#6b7280",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(107, 114, 128, 0))",
	},
	"from-gray-600": {
		from:  "#4b5563",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(75, 85, 99, 0))",
	},
	"from-gray-700": {
		from:  "#374151",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(55, 65, 81, 0))",
	},
	"from-gray-800": {
		from:  "#1f2937",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(31, 41, 55, 0))",
	},
	"from-gray-900": {
		from:  "#111827",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(17, 24, 39, 0))",
	},
	"from-red-50": {
		from:  "#fef2f2",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(254, 242, 242, 0))",
	},
	"from-red-100": {
		from:  "#fee2e2",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(254, 226, 226, 0))",
	},
	"from-red-200": {
		from:  "#fecaca",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(254, 202, 202, 0))",
	},
	"from-red-300": {
		from:  "#fca5a5",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(252, 165, 165, 0))",
	},
	"from-red-400": {
		from:  "#f87171",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(248, 113, 113, 0))",
	},
	"from-red-500": {
		from:  "#ef4444",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(239, 68, 68, 0))",
	},
	"from-red-600": {
		from:  "#dc2626",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(220, 38, 38, 0))",
	},
	"from-red-700": {
		from:  "#b91c1c",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(185, 28, 28, 0))",
	},
	"from-red-800": {
		from:  "#991b1b",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(153, 27, 27, 0))",
	},
	"from-red-900": {
		from:  "#7f1d1d",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(127, 29, 29, 0))",
	},
	"from-yellow-50": {
		from:  "#fffbeb",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(255, 251, 235, 0))",
	},
	"from-yellow-100": {
		from:  "#fef3c7",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(254, 243, 199, 0))",
	},
	"from-yellow-200": {
		from:  "#fde68a",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(253, 230, 138, 0))",
	},
	"from-yellow-300": {
		from:  "#fcd34d",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(252, 211, 77, 0))",
	},
	"from-yellow-400": {
		from:  "#fbbf24",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(251, 191, 36, 0))",
	},
	"from-yellow-500": {
		from:  "#f59e0b",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(245, 158, 11, 0))",
	},
	"from-yellow-600": {
		from:  "#d97706",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(217, 119, 6, 0))",
	},
	"from-yellow-700": {
		from:  "#b45309",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(180, 83, 9, 0))",
	},
	"from-yellow-800": {
		from:  "#92400e",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(146, 64, 14, 0))",
	},
	"from-yellow-900": {
		from:  "#78350f",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(120, 53, 15, 0))",
	},
	"from-green-50": {
		from:  "#ecfdf5",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(236, 253, 245, 0))",
	},
	"from-green-100": {
		from:  "#d1fae5",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(209, 250, 229, 0))",
	},
	"from-green-200": {
		from:  "#a7f3d0",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(167, 243, 208, 0))",
	},
	"from-green-300": {
		from:  "#6ee7b7",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(110, 231, 183, 0))",
	},
	"from-green-400": {
		from:  "#34d399",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(52, 211, 153, 0))",
	},
	"from-green-500": {
		from:  "#10b981",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(16, 185, 129, 0))",
	},
	"from-green-600": {
		from:  "#059669",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(5, 150, 105, 0))",
	},
	"from-green-700": {
		from:  "#047857",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(4, 120, 87, 0))",
	},
	"from-green-800": {
		from:  "#065f46",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(6, 95, 70, 0))",
	},
	"from-green-900": {
		from:  "#064e3b",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(6, 78, 59, 0))",
	},
	"from-blue-50": {
		from:  "#eff6ff",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(239, 246, 255, 0))",
	},
	"from-blue-100": {
		from:  "#dbeafe",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(219, 234, 254, 0))",
	},
	"from-blue-200": {
		from:  "#bfdbfe",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(191, 219, 254, 0))",
	},
	"from-blue-300": {
		from:  "#93c5fd",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(147, 197, 253, 0))",
	},
	"from-blue-400": {
		from:  "#60a5fa",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(96, 165, 250, 0))",
	},
	"from-blue-500": {
		from:  "#3b82f6",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(59, 130, 246, 0))",
	},
	"from-blue-600": {
		from:  "#2563eb",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(37, 99, 235, 0))",
	},
	"from-blue-700": {
		from:  "#1d4ed8",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(29, 78, 216, 0))",
	},
	"from-blue-800": {
		from:  "#1e40af",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(30, 64, 175, 0))",
	},
	"from-blue-900": {
		from:  "#1e3a8a",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(30, 58, 138, 0))",
	},
	"from-indigo-50": {
		from:  "#eef2ff",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(238, 242, 255, 0))",
	},
	"from-indigo-100": {
		from:  "#e0e7ff",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(224, 231, 255, 0))",
	},
	"from-indigo-200": {
		from:  "#c7d2fe",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(199, 210, 254, 0))",
	},
	"from-indigo-300": {
		from:  "#a5b4fc",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(165, 180, 252, 0))",
	},
	"from-indigo-400": {
		from:  "#818cf8",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(129, 140, 248, 0))",
	},
	"from-indigo-500": {
		from:  "#6366f1",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(99, 102, 241, 0))",
	},
	"from-indigo-600": {
		from:  "#4f46e5",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(79, 70, 229, 0))",
	},
	"from-indigo-700": {
		from:  "#4338ca",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(67, 56, 202, 0))",
	},
	"from-indigo-800": {
		from:  "#3730a3",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(55, 48, 163, 0))",
	},
	"from-indigo-900": {
		from:  "#312e81",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(49, 46, 129, 0))",
	},
	"from-purple-50": {
		from:  "#f5f3ff",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(245, 243, 255, 0))",
	},
	"from-purple-100": {
		from:  "#ede9fe",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(237, 233, 254, 0))",
	},
	"from-purple-200": {
		from:  "#ddd6fe",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(221, 214, 254, 0))",
	},
	"from-purple-300": {
		from:  "#c4b5fd",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(196, 181, 253, 0))",
	},
	"from-purple-400": {
		from:  "#a78bfa",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(167, 139, 250, 0))",
	},
	"from-purple-500": {
		from:  "#8b5cf6",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(139, 92, 246, 0))",
	},
	"from-purple-600": {
		from:  "#7c3aed",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(124, 58, 237, 0))",
	},
	"from-purple-700": {
		from:  "#6d28d9",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(109, 40, 217, 0))",
	},
	"from-purple-800": {
		from:  "#5b21b6",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(91, 33, 182, 0))",
	},
	"from-purple-900": {
		from:  "#4c1d95",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(76, 29, 149, 0))",
	},
	"from-pink-50": {
		from:  "#fdf2f8",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(253, 242, 248, 0))",
	},
	"from-pink-100": {
		from:  "#fce7f3",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(252, 231, 243, 0))",
	},
	"from-pink-200": {
		from:  "#fbcfe8",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(251, 207, 232, 0))",
	},
	"from-pink-300": {
		from:  "#f9a8d4",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(249, 168, 212, 0))",
	},
	"from-pink-400": {
		from:  "#f472b6",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(244, 114, 182, 0))",
	},
	"from-pink-500": {
		from:  "#ec4899",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(236, 72, 153, 0))",
	},
	"from-pink-600": {
		from:  "#db2777",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(219, 39, 119, 0))",
	},
	"from-pink-700": {
		from:  "#be185d",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(190, 24, 93, 0))",
	},
	"from-pink-800": {
		from:  "#9d174d",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(157, 23, 77, 0))",
	},
	"from-pink-900": {
		from:  "#831843",
		stops: "var(--tc-gradient-from), var(--tc-gradient-to, rgba(131, 24, 67, 0))",
	},
	"via-transparent": {
		stops: "var(--tc-gradient-from), transparent, var(--tc-gradient-to, rgba(0, 0, 0, 0))",
	},
	"via-current": {
		stops: "var(--tc-gradient-from), currentColor, var(--tc-gradient-to, rgba(255, 255, 255, 0))",
	},
	"via-black": {
		stops: "var(--tc-gradient-from), #000, var(--tc-gradient-to, rgba(0, 0, 0, 0))",
	},
	"via-white": {
		stops: "var(--tc-gradient-from), #fff, var(--tc-gradient-to, rgba(255, 255, 255, 0))",
	},
	"via-gray-50": {
		stops: "var(--tc-gradient-from), #f9fafb, var(--tc-gradient-to, rgba(249, 250, 251, 0))",
	},
	"via-gray-100": {
		stops: "var(--tc-gradient-from), #f3f4f6, var(--tc-gradient-to, rgba(243, 244, 246, 0))",
	},
	"via-gray-200": {
		stops: "var(--tc-gradient-from), #e5e7eb, var(--tc-gradient-to, rgba(229, 231, 235, 0))",
	},
	"via-gray-300": {
		stops: "var(--tc-gradient-from), #d1d5db, var(--tc-gradient-to, rgba(209, 213, 219, 0))",
	},
	"via-gray-400": {
		stops: "var(--tc-gradient-from), #9ca3af, var(--tc-gradient-to, rgba(156, 163, 175, 0))",
	},
	"via-gray-500": {
		stops: "var(--tc-gradient-from), #6b7280, var(--tc-gradient-to, rgba(107, 114, 128, 0))",
	},
	"via-gray-600": {
		stops: "var(--tc-gradient-from), #4b5563, var(--tc-gradient-to, rgba(75, 85, 99, 0))",
	},
	"via-gray-700": {
		stops: "var(--tc-gradient-from), #374151, var(--tc-gradient-to, rgba(55, 65, 81, 0))",
	},
	"via-gray-800": {
		stops: "var(--tc-gradient-from), #1f2937, var(--tc-gradient-to, rgba(31, 41, 55, 0))",
	},
	"via-gray-900": {
		stops: "var(--tc-gradient-from), #111827, var(--tc-gradient-to, rgba(17, 24, 39, 0))",
	},
	"via-red-50": {
		stops: "var(--tc-gradient-from), #fef2f2, var(--tc-gradient-to, rgba(254, 242, 242, 0))",
	},
	"via-red-100": {
		stops: "var(--tc-gradient-from), #fee2e2, var(--tc-gradient-to, rgba(254, 226, 226, 0))",
	},
	"via-red-200": {
		stops: "var(--tc-gradient-from), #fecaca, var(--tc-gradient-to, rgba(254, 202, 202, 0))",
	},
	"via-red-300": {
		stops: "var(--tc-gradient-from), #fca5a5, var(--tc-gradient-to, rgba(252, 165, 165, 0))",
	},
	"via-red-400": {
		stops: "var(--tc-gradient-from), #f87171, var(--tc-gradient-to, rgba(248, 113, 113, 0))",
	},
	"via-red-500": {
		stops: "var(--tc-gradient-from), #ef4444, var(--tc-gradient-to, rgba(239, 68, 68, 0))",
	},
	"via-red-600": {
		stops: "var(--tc-gradient-from), #dc2626, var(--tc-gradient-to, rgba(220, 38, 38, 0))",
	},
	"via-red-700": {
		stops: "var(--tc-gradient-from), #b91c1c, var(--tc-gradient-to, rgba(185, 28, 28, 0))",
	},
	"via-red-800": {
		stops: "var(--tc-gradient-from), #991b1b, var(--tc-gradient-to, rgba(153, 27, 27, 0))",
	},
	"via-red-900": {
		stops: "var(--tc-gradient-from), #7f1d1d, var(--tc-gradient-to, rgba(127, 29, 29, 0))",
	},
	"via-yellow-50": {
		stops: "var(--tc-gradient-from), #fffbeb, var(--tc-gradient-to, rgba(255, 251, 235, 0))",
	},
	"via-yellow-100": {
		stops: "var(--tc-gradient-from), #fef3c7, var(--tc-gradient-to, rgba(254, 243, 199, 0))",
	},
	"via-yellow-200": {
		stops: "var(--tc-gradient-from), #fde68a, var(--tc-gradient-to, rgba(253, 230, 138, 0))",
	},
	"via-yellow-300": {
		stops: "var(--tc-gradient-from), #fcd34d, var(--tc-gradient-to, rgba(252, 211, 77, 0))",
	},
	"via-yellow-400": {
		stops: "var(--tc-gradient-from), #fbbf24, var(--tc-gradient-to, rgba(251, 191, 36, 0))",
	},
	"via-yellow-500": {
		stops: "var(--tc-gradient-from), #f59e0b, var(--tc-gradient-to, rgba(245, 158, 11, 0))",
	},
	"via-yellow-600": {
		stops: "var(--tc-gradient-from), #d97706, var(--tc-gradient-to, rgba(217, 119, 6, 0))",
	},
	"via-yellow-700": {
		stops: "var(--tc-gradient-from), #b45309, var(--tc-gradient-to, rgba(180, 83, 9, 0))",
	},
	"via-yellow-800": {
		stops: "var(--tc-gradient-from), #92400e, var(--tc-gradient-to, rgba(146, 64, 14, 0))",
	},
	"via-yellow-900": {
		stops: "var(--tc-gradient-from), #78350f, var(--tc-gradient-to, rgba(120, 53, 15, 0))",
	},
	"via-green-50": {
		stops: "var(--tc-gradient-from), #ecfdf5, var(--tc-gradient-to, rgba(236, 253, 245, 0))",
	},
	"via-green-100": {
		stops: "var(--tc-gradient-from), #d1fae5, var(--tc-gradient-to, rgba(209, 250, 229, 0))",
	},
	"via-green-200": {
		stops: "var(--tc-gradient-from), #a7f3d0, var(--tc-gradient-to, rgba(167, 243, 208, 0))",
	},
	"via-green-300": {
		stops: "var(--tc-gradient-from), #6ee7b7, var(--tc-gradient-to, rgba(110, 231, 183, 0))",
	},
	"via-green-400": {
		stops: "var(--tc-gradient-from), #34d399, var(--tc-gradient-to, rgba(52, 211, 153, 0))",
	},
	"via-green-500": {
		stops: "var(--tc-gradient-from), #10b981, var(--tc-gradient-to, rgba(16, 185, 129, 0))",
	},
	"via-green-600": {
		stops: "var(--tc-gradient-from), #059669, var(--tc-gradient-to, rgba(5, 150, 105, 0))",
	},
	"via-green-700": {
		stops: "var(--tc-gradient-from), #047857, var(--tc-gradient-to, rgba(4, 120, 87, 0))",
	},
	"via-green-800": {
		stops: "var(--tc-gradient-from), #065f46, var(--tc-gradient-to, rgba(6, 95, 70, 0))",
	},
	"via-green-900": {
		stops: "var(--tc-gradient-from), #064e3b, var(--tc-gradient-to, rgba(6, 78, 59, 0))",
	},
	"via-blue-50": {
		stops: "var(--tc-gradient-from), #eff6ff, var(--tc-gradient-to, rgba(239, 246, 255, 0))",
	},
	"via-blue-100": {
		stops: "var(--tc-gradient-from), #dbeafe, var(--tc-gradient-to, rgba(219, 234, 254, 0))",
	},
	"via-blue-200": {
		stops: "var(--tc-gradient-from), #bfdbfe, var(--tc-gradient-to, rgba(191, 219, 254, 0))",
	},
	"via-blue-300": {
		stops: "var(--tc-gradient-from), #93c5fd, var(--tc-gradient-to, rgba(147, 197, 253, 0))",
	},
	"via-blue-400": {
		stops: "var(--tc-gradient-from), #60a5fa, var(--tc-gradient-to, rgba(96, 165, 250, 0))",
	},
	"via-blue-500": {
		stops: "var(--tc-gradient-from), #3b82f6, var(--tc-gradient-to, rgba(59, 130, 246, 0))",
	},
	"via-blue-600": {
		stops: "var(--tc-gradient-from), #2563eb, var(--tc-gradient-to, rgba(37, 99, 235, 0))",
	},
	"via-blue-700": {
		stops: "var(--tc-gradient-from), #1d4ed8, var(--tc-gradient-to, rgba(29, 78, 216, 0))",
	},
	"via-blue-800": {
		stops: "var(--tc-gradient-from), #1e40af, var(--tc-gradient-to, rgba(30, 64, 175, 0))",
	},
	"via-blue-900": {
		stops: "var(--tc-gradient-from), #1e3a8a, var(--tc-gradient-to, rgba(30, 58, 138, 0))",
	},
	"via-indigo-50": {
		stops: "var(--tc-gradient-from), #eef2ff, var(--tc-gradient-to, rgba(238, 242, 255, 0))",
	},
	"via-indigo-100": {
		stops: "var(--tc-gradient-from), #e0e7ff, var(--tc-gradient-to, rgba(224, 231, 255, 0))",
	},
	"via-indigo-200": {
		stops: "var(--tc-gradient-from), #c7d2fe, var(--tc-gradient-to, rgba(199, 210, 254, 0))",
	},
	"via-indigo-300": {
		stops: "var(--tc-gradient-from), #a5b4fc, var(--tc-gradient-to, rgba(165, 180, 252, 0))",
	},
	"via-indigo-400": {
		stops: "var(--tc-gradient-from), #818cf8, var(--tc-gradient-to, rgba(129, 140, 248, 0))",
	},
	"via-indigo-500": {
		stops: "var(--tc-gradient-from), #6366f1, var(--tc-gradient-to, rgba(99, 102, 241, 0))",
	},
	"via-indigo-600": {
		stops: "var(--tc-gradient-from), #4f46e5, var(--tc-gradient-to, rgba(79, 70, 229, 0))",
	},
	"via-indigo-700": {
		stops: "var(--tc-gradient-from), #4338ca, var(--tc-gradient-to, rgba(67, 56, 202, 0))",
	},
	"via-indigo-800": {
		stops: "var(--tc-gradient-from), #3730a3, var(--tc-gradient-to, rgba(55, 48, 163, 0))",
	},
	"via-indigo-900": {
		stops: "var(--tc-gradient-from), #312e81, var(--tc-gradient-to, rgba(49, 46, 129, 0))",
	},
	"via-purple-50": {
		stops: "var(--tc-gradient-from), #f5f3ff, var(--tc-gradient-to, rgba(245, 243, 255, 0))",
	},
	"via-purple-100": {
		stops: "var(--tc-gradient-from), #ede9fe, var(--tc-gradient-to, rgba(237, 233, 254, 0))",
	},
	"via-purple-200": {
		stops: "var(--tc-gradient-from), #ddd6fe, var(--tc-gradient-to, rgba(221, 214, 254, 0))",
	},
	"via-purple-300": {
		stops: "var(--tc-gradient-from), #c4b5fd, var(--tc-gradient-to, rgba(196, 181, 253, 0))",
	},
	"via-purple-400": {
		stops: "var(--tc-gradient-from), #a78bfa, var(--tc-gradient-to, rgba(167, 139, 250, 0))",
	},
	"via-purple-500": {
		stops: "var(--tc-gradient-from), #8b5cf6, var(--tc-gradient-to, rgba(139, 92, 246, 0))",
	},
	"via-purple-600": {
		stops: "var(--tc-gradient-from), #7c3aed, var(--tc-gradient-to, rgba(124, 58, 237, 0))",
	},
	"via-purple-700": {
		stops: "var(--tc-gradient-from), #6d28d9, var(--tc-gradient-to, rgba(109, 40, 217, 0))",
	},
	"via-purple-800": {
		stops: "var(--tc-gradient-from), #5b21b6, var(--tc-gradient-to, rgba(91, 33, 182, 0))",
	},
	"via-purple-900": {
		stops: "var(--tc-gradient-from), #4c1d95, var(--tc-gradient-to, rgba(76, 29, 149, 0))",
	},
	"via-pink-50": {
		stops: "var(--tc-gradient-from), #fdf2f8, var(--tc-gradient-to, rgba(253, 242, 248, 0))",
	},
	"via-pink-100": {
		stops: "var(--tc-gradient-from), #fce7f3, var(--tc-gradient-to, rgba(252, 231, 243, 0))",
	},
	"via-pink-200": {
		stops: "var(--tc-gradient-from), #fbcfe8, var(--tc-gradient-to, rgba(251, 207, 232, 0))",
	},
	"via-pink-300": {
		stops: "var(--tc-gradient-from), #f9a8d4, var(--tc-gradient-to, rgba(249, 168, 212, 0))",
	},
	"via-pink-400": {
		stops: "var(--tc-gradient-from), #f472b6, var(--tc-gradient-to, rgba(244, 114, 182, 0))",
	},
	"via-pink-500": {
		stops: "var(--tc-gradient-from), #ec4899, var(--tc-gradient-to, rgba(236, 72, 153, 0))",
	},
	"via-pink-600": {
		stops: "var(--tc-gradient-from), #db2777, var(--tc-gradient-to, rgba(219, 39, 119, 0))",
	},
	"via-pink-700": {
		stops: "var(--tc-gradient-from), #be185d, var(--tc-gradient-to, rgba(190, 24, 93, 0))",
	},
	"via-pink-800": {
		stops: "var(--tc-gradient-from), #9d174d, var(--tc-gradient-to, rgba(157, 23, 77, 0))",
	},
	"via-pink-900": {
		stops: "var(--tc-gradient-from), #831843, var(--tc-gradient-to, rgba(131, 24, 67, 0))",
	},
	"to-transparent": {
		to: "transparent",
	},
	"to-current": {
		to: "currentColor",
	},
	"to-black": {
		to: "#000",
	},
	"to-white": {
		to: "#fff",
	},
	"to-gray-50": {
		to: "#f9fafb",
	},
	"to-gray-100": {
		to: "#f3f4f6",
	},
	"to-gray-200": {
		to: "#e5e7eb",
	},
	"to-gray-300": {
		to: "#d1d5db",
	},
	"to-gray-400": {
		to: "#9ca3af",
	},
	"to-gray-500": {
		to: "#6b7280",
	},
	"to-gray-600": {
		to: "#4b5563",
	},
	"to-gray-700": {
		to: "#374151",
	},
	"to-gray-800": {
		to: "#1f2937",
	},
	"to-gray-900": {
		to: "#111827",
	},
	"to-red-50": {
		to: "#fef2f2",
	},
	"to-red-100": {
		to: "#fee2e2",
	},
	"to-red-200": {
		to: "#fecaca",
	},
	"to-red-300": {
		to: "#fca5a5",
	},
	"to-red-400": {
		to: "#f87171",
	},
	"to-red-500": {
		to: "#ef4444",
	},
	"to-red-600": {
		to: "#dc2626",
	},
	"to-red-700": {
		to: "#b91c1c",
	},
	"to-red-800": {
		to: "#991b1b",
	},
	"to-red-900": {
		to: "#7f1d1d",
	},
	"to-yellow-50": {
		to: "#fffbeb",
	},
	"to-yellow-100": {
		to: "#fef3c7",
	},
	"to-yellow-200": {
		to: "#fde68a",
	},
	"to-yellow-300": {
		to: "#fcd34d",
	},
	"to-yellow-400": {
		to: "#fbbf24",
	},
	"to-yellow-500": {
		to: "#f59e0b",
	},
	"to-yellow-600": {
		to: "#d97706",
	},
	"to-yellow-700": {
		to: "#b45309",
	},
	"to-yellow-800": {
		to: "#92400e",
	},
	"to-yellow-900": {
		to: "#78350f",
	},
	"to-green-50": {
		to: "#ecfdf5",
	},
	"to-green-100": {
		to: "#d1fae5",
	},
	"to-green-200": {
		to: "#a7f3d0",
	},
	"to-green-300": {
		to: "#6ee7b7",
	},
	"to-green-400": {
		to: "#34d399",
	},
	"to-green-500": {
		to: "#10b981",
	},
	"to-green-600": {
		to: "#059669",
	},
	"to-green-700": {
		to: "#047857",
	},
	"to-green-800": {
		to: "#065f46",
	},
	"to-green-900": {
		to: "#064e3b",
	},
	"to-blue-50": {
		to: "#eff6ff",
	},
	"to-blue-100": {
		to: "#dbeafe",
	},
	"to-blue-200": {
		to: "#bfdbfe",
	},
	"to-blue-300": {
		to: "#93c5fd",
	},
	"to-blue-400": {
		to: "#60a5fa",
	},
	"to-blue-500": {
		to: "#3b82f6",
	},
	"to-blue-600": {
		to: "#2563eb",
	},
	"to-blue-700": {
		to: "#1d4ed8",
	},
	"to-blue-800": {
		to: "#1e40af",
	},
	"to-blue-900": {
		to: "#1e3a8a",
	},
	"to-indigo-50": {
		to: "#eef2ff",
	},
	"to-indigo-100": {
		to: "#e0e7ff",
	},
	"to-indigo-200": {
		to: "#c7d2fe",
	},
	"to-indigo-300": {
		to: "#a5b4fc",
	},
	"to-indigo-400": {
		to: "#818cf8",
	},
	"to-indigo-500": {
		to: "#6366f1",
	},
	"to-indigo-600": {
		to: "#4f46e5",
	},
	"to-indigo-700": {
		to: "#4338ca",
	},
	"to-indigo-800": {
		to: "#3730a3",
	},
	"to-indigo-900": {
		to: "#312e81",
	},
	"to-purple-50": {
		to: "#f5f3ff",
	},
	"to-purple-100": {
		to: "#ede9fe",
	},
	"to-purple-200": {
		to: "#ddd6fe",
	},
	"to-purple-300": {
		to: "#c4b5fd",
	},
	"to-purple-400": {
		to: "#a78bfa",
	},
	"to-purple-500": {
		to: "#8b5cf6",
	},
	"to-purple-600": {
		to: "#7c3aed",
	},
	"to-purple-700": {
		to: "#6d28d9",
	},
	"to-purple-800": {
		to: "#5b21b6",
	},
	"to-purple-900": {
		to: "#4c1d95",
	},
	"to-pink-50": {
		to: "#fdf2f8",
	},
	"to-pink-100": {
		to: "#fce7f3",
	},
	"to-pink-200": {
		to: "#fbcfe8",
	},
	"to-pink-300": {
		to: "#f9a8d4",
	},
	"to-pink-400": {
		to: "#f472b6",
	},
	"to-pink-500": {
		to: "#ec4899",
	},
	"to-pink-600": {
		to: "#db2777",
	},
	"to-pink-700": {
		to: "#be185d",
	},
	"to-pink-800": {
		to: "#9d174d",
	},
	"to-pink-900": {
		to: "#831843",
	},
}

func (cg *CSSBuilder) GenerateGradients(prefixes ...string) {
	prefix := GetPrefix(prefixes...)
	className := ""
	for class, g := range gradientColors {
		if prefix != "" {
			className = fmt.Sprintf(".%s\\:%s", prefix, class)
		} else {
			className = "." + class
		}

		if g.from != "" {
			fmt.Fprintf(cg.Output, "%s {\n --tc-gradient-from:: %s;\n --tc-gradient-stops: %s;\n}\n\n", className, g.from, g.stops)
			continue
		} else if g.stops != "" {
			fmt.Fprintf(cg.Output, "%s {\n --tc-gradient-stops: %s;\n}\n\n", className, g.stops)
			continue
		}

		fmt.Fprintf(cg.Output, "%s {\n --tc-gradient-to: %s;\n}\n\n", className, g.to)
	}
}
