#include <stdio.h>
#include <stdarg.h>

// Startup/shutdown
void Niall_Init(void);
void Niall_Free(void);

// Niall's main functions
void Niall_Learn(char *Buffer);
void Niall_Reply(char *Buffer,int BufSize);

// Housekeeping functions
void Niall_NewDictionary(void);
void Niall_ListDictionary(void);
void Niall_SaveDictionary(char *file);
void Niall_LoadDictionary(char *file);
void Niall_CorrectSpelling(char *Original,char *Correct);

// Niall calls these functions (from C to Go)
extern void niall_go_print( char *msg );
extern void niall_go_warning( char *msg );
extern void niall_go_error( char *msg );

