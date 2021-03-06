//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package jbig2 ;import (_g "github.com/unidoc/unipdf/v3/internal/bitwise";_d "github.com/unidoc/unipdf/v3/internal/jbig2/decoder";_df "github.com/unidoc/unipdf/v3/internal/jbig2/document";_c "github.com/unidoc/unipdf/v3/internal/jbig2/document/segments";_da "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_b "sort";);type Globals map[int ]*_c .Header ;func DecodeBytes (encoded []byte ,parameters _d .Parameters ,globals ...Globals )([]byte ,error ){var _bd Globals ;if len (globals )> 0{_bd =globals [0];};_e ,_gd :=_d .Decode (encoded ,parameters ,_bd .ToDocumentGlobals ());if _gd !=nil {return nil ,_gd ;};return _e .DecodeNextPage ();};func DecodeGlobals (encoded []byte )(Globals ,error ){const _dc ="\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073";_a :=_g .NewReader (encoded );_bdd ,_ed :=_df .DecodeDocument (_a ,nil );if _ed !=nil {return nil ,_da .Wrap (_ed ,_dc ,"");};if _bdd .GlobalSegments ==nil ||(_bdd .GlobalSegments .Segments ==nil ){return nil ,_da .Error (_dc ,"\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};_ec :=Globals {};for _ ,_de :=range _bdd .GlobalSegments .Segments {_ec [int (_de .SegmentNumber )]=_de ;};return _ec ,nil ;};func (_af Globals )ToDocumentGlobals ()*_df .Globals {if _af ==nil {return nil ;};_eb :=[]*_c .Header {};for _ ,_ee :=range _af {_eb =append (_eb ,_ee );};_b .Slice (_eb ,func (_ae ,_bg int )bool {return _eb [_ae ].SegmentNumber < _eb [_bg ].SegmentNumber });return &_df .Globals {Segments :_eb };};