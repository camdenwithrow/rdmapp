"use strict";

function darkenBackground(background) {
	background.classList.remove("bg-zinc-50/0")
	background.classList.add("bg-zinc-900/80")
	document.documentElement.classList.add("bg-zinc-900/80")
}

function removeDarkenedBackground(background) {
	background.classList.add("bg-zinc-50/0")
	background.classList.remove("bg-zinc-900/80")
	document.documentElement.classList.remove("bg-zinc-900/80")
}

function getSheetEls(id) {
	const container = document.getElementById(`${id}-sheet-container`)
	const background = document.getElementById(`${id}-sheet-bg`)
	const sheet = document.getElementById(`${id}-sheet`)
	return [container, background, sheet]
}

function getDialogEls(id) {
	console.log(`${id}-dialog-container`)
	const container = document.getElementById(`${id}-dialog-container`)
	const background = document.getElementById(`${id}-dialog-bg`)
	const dialog = document.getElementById(`${id}-dialog`)
	return [container, background, dialog]
}

function openSheet(id) {
	const [container, background, sheet] = getSheetEls(id)

	container.classList.remove("hidden")

	setTimeout(() => {
		darkenBackground(background)

		sheet.classList.remove("translate-x-full")
		sheet.classList.add("translate-x-0")
	}, 1)
}

function closeSheet(id) {
	const [container, background, sheet] = getSheetEls(id)

	removeDarkenedBackground(background)

	sheet.classList.add("translate-x-full")
	sheet.classList.remove("translate-x-0")

	setTimeout(() => {
		container.classList.add("hidden")
	}, 400)
}

function openDialog(id) {
	const [container, background, dialog] = getDialogEls(id)

	container.classList.remove("hidden")

	setTimeout(() => {
		darkenBackground(background)

		dialog.classList.remove("scale-90", "opacity-0")
		dialog.classList.add("scale-100", "opacity-100")
	}, 1)
}

function closeDialog(id) {
	const [container, background, dialog] = getDialogEls(id)

	removeDarkenedBackground(background)

	dialog.classList.add("scale-90", "opacity-0")
	dialog.classList.remove("scale-100", "opacity-100")

	setTimeout(() => {
		container.classList.add("hidden")
	}, 400)
}
