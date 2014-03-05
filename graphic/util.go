/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package graphic

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func openFileOrCreate(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
}

func encodeImage(w io.Writer, m image.Image, f format) (err error) {
	switch f {
	case PNG:
		err = png.Encode(w, m)
	case JPEG:
		err = jpeg.Encode(w, m, nil)
	default:
		err = png.Encode(w, m)
	}
	return
}
