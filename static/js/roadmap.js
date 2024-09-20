"use strict";

function getSheetEls(id) {
	const container = document.getElementById(`${id}-sheet-container`)
	const background = document.getElementById(`${id}-sheet-bg`)
	const sheet = document.getElementById(`${id}-sheet`)
	return [container, background, sheet]
}

function openSheet(id) {
	const [container, background, sheet] = getSheetEls(id)
	console.log(`${id}-sheet-container`)

	container.classList.remove("hidden")

	setTimeout(() => {
		background.classList.remove("bg-zinc-50/0")
		background.classList.add("bg-zinc-900/80")
		document.documentElement.classList.add("bg-zinc-900/80")

		sheet.classList.remove("translate-x-full")
		sheet.classList.add("translate-x-0")
	}, 1)
}

function closeSheet(id) {
	const [container, background, sheet] = getSheetEls(id)

	background.classList.add("bg-zinc-50/0")
	background.classList.remove("bg-zinc-900/80")
	document.documentElement.classList.remove("bg-zinc-900/80")

	sheet.classList.add("translate-x-full")
	sheet.classList.remove("translate-x-0")

	setTimeout(() => {
		container.classList.add("hidden")
	}, 500)
}
