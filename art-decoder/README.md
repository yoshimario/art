# Art Decoder - Command Line Tool

The **Art Decoder** is a command-line tool designed to help artists like Chris generate text-based art more efficiently. It allows users to encode and decode text-based art using a simple syntax for repeating characters. The tool also includes additional features like multi-line decoding, encoding mode, animations, and sound effects to enhance the user experience.

## Features

- **Decoder**: Converts encoded strings into text-based art.
- **Encoder**: Converts text-based art into encoded strings.
- **Multi-line Decoding**: Supports decoding of multi-line encoded strings.
- **Animations**: Adds visual effects like typing, rainbow, banner, and loading animations.
- **Sound Effects**: Plays sound effects like beeps, success tones, and error sounds.
- **Error Handling**: Provides clear error messages for malformed inputs.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/art-decoder.git
   cd art-decoder
   ```

2. **Build the tool**:

   ```bash
   go build -o art-decoder
   ```

3. **Run the tool**:

   ```bash
   go run . [options] "[encoded string]"
   ```

## Usage

### Basic Decoding

To decode a single-line encoded string:

```bash
go run . "[5 #][5 -_]-[5 #]"
```

**Output**:

```bash
#####-_-_-_-_-_-#####
```

### Multi-line Decoding

To decode a multi-line encoded string, use the `-ml` flag:

```bash
go run . -ml "[5 #][5 -_]-[5 #]\n[3 @][2 !]"
```

**Output**:

```
#####-_-_-_-_-_-#####
@@@!!
```

### Encoding Mode

To encode text-based art into the encoded format, use the `-encode` flag:

```bash
go run . -encode "#####-_-_-_-_-_-#####"
```

**Output**:

```
[5 #][5 -_]-[5 #]
```

### Animations

Enable animations with the `-animate` flag and specify the animation type with `-animation-type`:

```bash
go run . -animate -animation-type=typing "[5 #][5 -_]-[5 #]"
```

**Available Animation Types**:

- `typing`: Typing effect.
- `rainbow`: Rainbow color effect.
- `banner`: Scrolling banner effect.
- `loading`: Loading bar animation.

### Sound Effects

Enable sound effects with the `-sound` flag and specify the sound type with `-sound-type`:

```bash
go run . -sound -sound-type=typing "[5 #][5 -_]-[5 #]"
```

**Available Sound Types**:

- `beep`: Simple beep sound.
- `success`: Success tone.
- `error`: Error sound.
- `typing`: Typing sound.
- `loading`: Loading sound.
- `banner`: Banner sound.

### Disable Colors

To disable colored output, use the `-no-color` flag:

```bash
go run . -no-color "[5 #][5 -_]-[5 #]"
```

### Adjust Animation Speed

Adjust the animation speed with the `-speed` flag (lower values are faster):

```bash
go run . -animate -animation-type=typing -speed=10 "[5 #][5 -_]-[5 #]"
```

### Adjust Sound Volume and Pitch

Adjust the sound volume and pitch with the `-volume` and `-pitch` flags:

```bash
go run . -sound -sound-type=typing -volume=75 -pitch=1.5 "[5 #][5 -_]-[5 #]"
```

## Error Handling

The tool provides clear error messages for malformed inputs. For example:

```
go run . "[5 #][5 -_-[5 #]"
```

**Output**:

```
Error: Missing closing bracket
```

## Examples

### Example 1: Decoding a Simple Pattern

```bash
go run . "[3 @][2 !]"
```

**Output**:

```
@@@!!
```

### Example 2: Decoding a Multi-line Pattern

```bash
go run . -ml "[5 #][5 -_]-[5 #]\n[3 @][2 !]"
```

**Output**:

```
#####-_-_-_-_-_-#####
@@@!!
```

### Example 3: Encoding a Text-based Art

```bash
go run . -encode "#####-_-_-_-_-_-#####"
```

**Output**:

```
[5 #][5 -_]-[5 #]
```

### Example 4: Using Animations and Sound Effects

```bash
go run . -animate -animation-type=typing -sound -sound-type=typing "[5 #][5 -_]-[5 #]"
```

**Output**:

```
#####-_-_-_-_-_-#####
```

*(With typing animation and sound effects)*

## Resources

- **Sample Files**:
  - `cats.encoded.txt`
  - `cats.art.txt`
  - `kood.encoded.txt`
  - `kood.art.txt`
  - `lion.encoded.txt`
  - `lion.art.txt`
  - `plane.encoded.txt`
  - `plane.art.txt`
- **Useful Links**:
  - [ASCII Art Archive](http://www.ascii-art.de/)
  - [Go Programming Language](https://golang.org/)

## What You'll Learn

- **Encoding**: Converting text-based art into a compressed format.
- **Parsing Functional Arguments**: Handling command-line arguments and flags.
- **Error Handling**: Providing clear error messages for malformed inputs.
- **Animations and Sound Effects**: Enhancing user experience with visual and auditory feedback.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://chat.deepseek.com/a/chat/s/LICENSE) file for details.

------

Enjoy creating text-based art with **Art Decoder**! 🎨