# Legend

<b>b</b>: 3-bit value.<p>

<b>n</b>: 8-bit value.<p>

<b>nn</b>: 16-bit value, <i>little-endian. E.g. the JP 0123H opcode is C3 23 01.<i><p>

<b>o</b>: 8-bit offset, <i>2’s complement.</i><p>

<b>r</b>: Register. This can be A, B, C, D, E, H or L. Add the following value to the last byte of the opcode<p>
Register |Register bits value
|:-: |:-:
B| 0
C| 1
D| 2
E| 3
H| 4
L| 5
A| 7

<b>IXp</b>: Denotes the high or low part of the IX register: IXh or IXl. Add the following value to the last byte of the opcode<p>
Register |Register bits value
|:-: |:-: 
IXh| 4
IXl| 5

<b>IYq</b>: Denotes the high or low part of the IY register: IYh or IYl. Add the following value to the last byte of the opcode<p>
Register |Register bits value
|:-: |:-:
IYh	|4
IYl	|5

<b>p</b>: Register where H and L have been replaced by IXh and IXl. Add the following value to the last byte of the opcode<p>
Register |Register bits value
|:-: |:-:
B|	0
C|	1
D|	2
E|	3
IXh|	4
IXl|	5
A|	7

<b>q</b>:Register where H and L have been replaced by IYh and IYl. Add the following value to the last byte of the opcode<p>
Register |Register bits value
|:-: |:-:
B|	0
C|	1
D|	2
E|	3
IYh| 4
IYl| 5
A|	7


ref: http://map.grauw.nl/resources/z80instr.php