Got it. I’ll expand the animations and transitions section to include:

* Syntax and examples of using rough markers like circles, highlights, underlines, and boxes in Slidev.
* How to animate these highlights with click interactions or transitions.
* Clear, simple recommendations on when and why to use each highlight type to emphasize ideas without overcomplicating the slide.

I’ll update you when that section is ready for review.

# Using Slidev’s Rough Markers and Click Animations for Emphasis

## Rough Markers for Visual Emphasis

&#x20;*Example of Slidev rough markers: the word **“underline”** is underlined in red, and **“highlight”** is marked with a yellow highlight.* Slidev integrates **Rough Notation** (a hand-drawn annotation library) via the `v-mark` directive to let you mark or highlight elements in your slides. In practice, you wrap the text you want to emphasize in a span with a `v-mark` directive. The format is typically `v-mark.<type>.<color>` – where **type** defines the style of marker (underline, circle, highlight, box, etc.) and **color** specifies an optional color theme. By default, if no type is given, it uses an underline style (and if no color is given, a default color is used). You can choose from built-in UnoCSS color names (e.g. `.red`, `.yellow`, `.blue`) or even supply a custom color via an object (e.g. `v-mark="{ color: '#234' }"` for a specific hex color).

* **Underline:** draws a sketchy underline beneath the text. For example, `<span v-mark.underline.red>Important</span>` will underline “Important” in red (you can omit `.underline` since it’s the default type). Use underlines to call out key terms or definitions within a sentence.
* **Highlight:** creates a rough highlighter effect behind the text (like swiping a marker pen). For example, `<span v-mark.highlight.yellow>Important</span>` wraps the word in a yellow highlight. This style is great for emphasizing a short phrase or result in a paragraph.
* **Circle:** draws a hand-drawn circle around the content. For example, `<span v-mark.circle.blue>Important</span>` will circle the word in a blue ink-style oval. Circles work well to spotlight a standalone item (like an icon, number, or a word) without altering the text background.
* **Box:** sketches a rough rectangle around the element. For example, `<span v-mark.box.green>Important</span>` will put a green rough-edged box around the text. A boxed highlight is useful for grouping a few words or drawing attention to a block like a code snippet or equation.

Each of these marker types appears as a hand-drawn annotation over your slide content, providing a visual emphasis that feels more engaging than a simple bold or italic. You can mix and match types and colors as needed – for instance, Slidev allows inline markers like `<span v-mark.underline.orange>inline markers</span>` (as shown in the official demo). Just be mindful of contrast: if you use a highlight or box, choose a color that keeps the text readable.

## Animating Markers with Click Interactions

One powerful feature of `v-mark` is that it doubles as an interactive click trigger. In fact, **`v-mark` behaves like `v-click` by default**, meaning the marker (and the text inside its span) will **not appear until you advance the slide** (click or press the next key). This allows you to reveal highlights step-by-step during your presentation. For example, if you have two marked phrases in a slide, the first will animate in on the first click, and the second on the next click (as seen in the code example where “underline” appears, then “highlight” appears on a subsequent click). Slidev’s click system assigns each `v-mark` element a sequential order automatically if no explicit order is given.

You can also **control the timing** of marker animations manually. The `v-mark` directive accepts the same modifiers as `v-click` for ordering:

* Use an absolute step index: e.g. `<span v-mark="3" v-mark.circle>Note</span>` will only show the circled “Note” on the 3rd click of that slide.
* Use a relative increment: e.g. `<span v-mark="+1" v-mark.highlight>Example</span>` will show that highlighted “Example” one click after the previous reveal.

When a marked element is triggered, the rough annotation is drawn with an animation (a short scribbling effect) to catch the eye. This happens in place during your slide show, without requiring any manual animation coding – Slidev handles it for you. If you navigate to a new slide, the click count resets for that slide (each slide has its own `$clicks` state), so markers on the next slide will start hidden again until triggered. **Slide transitions** (such as fade or slide-in effects between slides) do not interfere with marker animations – the markers will simply animate when their turn comes after the slide transition. Once revealed, the annotations remain on screen as part of the slide content. If you go back to a previous slide, by default Slidev will reset that slide’s clicks, allowing the marker animations to play again (so you can rehearse or revisit an explanation). In essence, rough markers integrate smoothly with Slidev’s navigation: they appear in sequence with other `v-click` elements, and they redraw on each viewing of the slide, making them reliable for live presentations.

## Best Practices for Using Highlights and Markers

Using rough markers can greatly enhance focus on key points, but it’s important to use them judiciously for a clear, professional look:

* **Choose the right marker style:** Match the marker type to the content. Use underlines for single keywords or terms (e.g. a vocabulary word or a variable name in code). Use a highlight for slightly longer phrases or important quotes – the background color makes them stand out. Use circles or boxes to encircle items like numbers, icons, or short phrases that are separate from a sentence (for example, circling a critical value on a chart, or boxing an equation result).
* **Limit the usage per slide:** It’s best to highlight only the most important point(s) on a slide. Overusing markers can clutter your slide and dilute their impact. One guideline is to *“limit underlining to important titles, names, or terms that require emphasis”* and **avoid underlining or highlighting entire sentences or large blocks of text**, as that can distract from the main content. In practice, try to stick to one or two emphasized elements per slide at most – make them count.
* **Maintain clarity and contrast:** Ensure the highlighted text remains readable. If you use a yellow highlight, for instance, the text should be dark (black or dark color) so it’s high contrast. Similarly, a red underline is best used under dark text. Slidev’s rough marks are semi-transparent by design, but always double-check that your audience can read the emphasized text easily, especially in a large room or on a projector.
* **Consistent styling:** For an academic or technical presentation, consistency in your emphasis helps the audience understand the meaning. Try not to use a rainbow of different colors for highlights without reason – pick one or two highlight colors that fit your theme (for example, maybe use yellow for general highlights and red for “warning/critical” points). Similarly, use the same marker type for similar kinds of information. For instance, if you underline a definition on one slide, also use underlines (not circles) for definitions on later slides. Consistency makes your slides feel cohesive and the emphasis appear intentional.
* **Use with spoken cues:** Plan your talk so that you trigger the marker animation at the moment you want the audience’s attention on that element. Rough markers are excellent for guiding focus during an explanation – e.g. *“As you can see here, the **results increased by 20%**”* (click to draw a circle around the “20%” on the slide). By timing the appearance of the highlight with your speech, you create a strong visual cue. After the point is made, move on so the audience isn’t fixated on the marker for too long. Remember that the novelty of the hand-drawn effect can draw attention, so use it to reinforce your message at the right time.

By following these practices, you can leverage Slidev’s animation and transition features to effectively emphasize important ideas without overwhelming your audience. Rough markers, when used sparingly and thoughtfully, add a dynamic, engaging touch to technical presentations while keeping the focus on the content. Each underline, highlight, circle, or box should answer the question: *“What do I want my audience to remember or notice first on this slide?”* Use them to make that answer unmistakably clear.

**Sources:** Slidev Documentation; Wim Deblauwe’s Slidev Tutorial; Rough Notation library docs; Stanford Snipe Markdown Guide (emphasis usage tips).
